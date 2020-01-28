package apihandler

import (
	RecommendationService "github.com/transavro/RecommenderService/proto"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strings"
	"time"
)

type RecommendationServiceHandler struct {
	RedisConnection  *redis.Client
	CloudwalkerRecommendation *CloudwalkerRedrec
	MongoCollection *mongo.Collection
}

func (h *RecommendationServiceHandler) TileClicked(ctx context.Context, req *RecommendationService.TileClickedRequest ) (*RecommendationService.TileClickedResponse, error) {
	go h.CloudwalkerRecommendation.Rate(req.GetTileId(), req.GetUserId(), req.GetTileScore())
	return &RecommendationService.TileClickedResponse{Success:true}, nil
}

func (h *RecommendationServiceHandler) GetCollabrativeFilteringData( req *RecommendationService.GetRecommendationRequest, stream RecommendationService.RecommendationService_GetCollabrativeFilteringDataServer) error{
	//err := h.CloudwalkerRecommendation.BatchUpdateSimilarUsers(2)
	//if err != nil {
	//	return err
	//}
	err := h.CloudwalkerRecommendation.UpdateSuggestedItems(req.UserId, 50)
	if err != nil {
		return err
	}
	userRedisKey := fmt.Sprintf("user:%s:collabrativeFilter", req.UserId)
	result, err :=  h.CheckInRedis(userRedisKey)
	if err != nil {
		return err
	}
	log.Println("Is Colab There ?", result)
	if result {
		stringResultCmd := h.RedisConnection.SMembers(userRedisKey)
		resultStringArray := stringResultCmd.Val()
		for i := 0 ; i < len(resultStringArray) ; i++ {
			var resultMovie RecommendationService.ContentTile
			err := proto.Unmarshal([]byte(resultStringArray[i]), &resultMovie)
			if err != nil {
				return err
			}
			err = stream.Send(&resultMovie)
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		resultSet, err := h.CloudwalkerRecommendation.GetUserSuggestions(req.UserId, 5)
		if err != nil {
			return err
		}
		if len(resultSet) > 0 {
			for _,v := range resultSet {
				result := h.RedisConnection.HGet("cloudwalkerTiles", v)
				if result.Err() != nil {
					return result.Err()
				}
				resultInBytes, _ := result.Bytes()
				h.RedisConnection.Del(userRedisKey)
				h.RedisConnection.SAdd(userRedisKey,resultInBytes)
				var movieTiles RecommendationService.ContentTile
				err := proto.Unmarshal(resultInBytes, &movieTiles)
				if err != nil {
					return err
				}
				err = stream.Send(&movieTiles)
				if err != nil {
					return err
				}
			}
			return nil
		}else {
			return status.Error(codes.Internal, fmt.Sprintf("No Recommendation"))
		}
	}
}

func (h *RecommendationServiceHandler) InitialRecommendationEngine(ctx context.Context, req *RecommendationService.InitRecommendationRequest ) (*RecommendationService.InitRecommendationResponse, error) {
	log.Println("Triggered ")

	go func(genres []string) {
		for _,v := range genres {
			genreKey :=  fmt.Sprintf("genre:%s:items", v)
			genreKey = formatString(genreKey)

			myStages := mongo.Pipeline{
				//stage 1
				bson.D{{"$match", bson.D{{"metadata.genre", bson.D{{"$in", bson.A{v}}}}}}},
				//stage 2
				bson.D{{"$sort", bson.D{{"created_at", -1}}}},
				//stage 3
				bson.D{{"$project", bson.D{{"_id", 0},{"ref_id", 1}}}},
			}
			cur, err := h.MongoCollection.Aggregate(context.Background(), myStages, options.Aggregate().SetMaxTime(2000*time.Millisecond))
			if err != nil {
				log.Println("error 1 ")
			}


			count := 0
			for cur.Next(ctx) {
				count++
				h.RedisConnection.SAdd(genreKey, cur.Current.Lookup("ref_id").StringValue())
			}
			log.Println(genreKey , " ==========>  ", count)
			cur.Close(ctx)
		}

	}(req.GetGenres())

	go func(languages []string) {
		// language Tiles Segregation
		for _,v := range languages {
			languageKey :=  fmt.Sprintf("languages:%s:items", v)
			languageKey = formatString(languageKey)
			myStages := mongo.Pipeline{
				//stage 1
				bson.D{{"$match", bson.D{{"metadata.languages", bson.D{{"$in", bson.A{v}}}}}}},
				//stage 2
				bson.D{{"$sort", bson.D{{"created_at", -1}}}},
				//stage 3
				bson.D{{"$project", bson.D{{"_id", 0},{"ref_id", 1}}}},
			}
			cur, err := h.MongoCollection.Aggregate(context.Background(), myStages, options.Aggregate().SetMaxTime(2000*time.Millisecond))
			if err != nil {
				log.Println("error 3 ")
			}

			count := 0
			for cur.Next(ctx) {
				count++
				h.RedisConnection.SAdd(languageKey, cur.Current.Lookup("ref_id").StringValue())
			}


			log.Println( languageKey, " ==========>  ",count)

			cur.Close(ctx)
		}
	}(req.GetLanguages())

	go func(categories []string) {
		for _,v := range req.Categories {
			categoriesKey :=  fmt.Sprintf("categories:%s:items", v)
			categoriesKey = formatString(categoriesKey)
			log.Println("categoriesKey  =====> ", categoriesKey)
			myStages := mongo.Pipeline{
				//stage 1
				bson.D{{"$match", bson.D{{"metadata.categories", bson.D{{"$in", bson.A{v}}}}}}},
				//stage 2
				bson.D{{"$sort", bson.D{{"created_at", -1}}}},
				//stage 3
				bson.D{{"$project", bson.D{{"_id", 0},{"ref_id", 1}}}},
			}
			cur, err := h.MongoCollection.Aggregate(context.Background(), myStages, options.Aggregate().SetMaxTime(2000*time.Millisecond))
			if err != nil {
				log.Println("error 5 ")

			}

			count := 0
			for cur.Next(ctx) {
				count++
				h.RedisConnection.SAdd(categoriesKey, cur.Current.Lookup("ref_id").StringValue())
			}


			log.Println( categoriesKey, " ==========>  ",count)

			cur.Close(ctx)
		}
	}(req.GetCategories())

	return &RecommendationService.InitRecommendationResponse{IsDone:true}, nil
}
// helper function
func formatString(value string) string {
	return strings.ToLower(strings.Replace(value, " ", "_", -1))
}

func (h *RecommendationServiceHandler) GetContentbasedData(req *RecommendationService.GetRecommendationRequest, stream RecommendationService.RecommendationService_GetContentbasedDataServer) error {
	isThere, err := h.CheckInRedis(fmt.Sprintf("user:%s:Contentbased", req.UserId))
	if err != nil {
		return err
	}
	log.Println("Is ContentBased There ? ", isThere)
	if isThere {
		// get the updated Tile Clicked By the User
		h.DiffContentBaseToTileClickedByUser(req.UserId)
		result := h.RedisConnection.SRandMemberN(fmt.Sprintf("user:%s:Contentbased", req.UserId), 30)
		if result.Err() != nil {
			return result.Err()
		}
		tilesResult := h.RedisConnection.HMGet("cloudwalkerTiles", result.Val()...)
		for i := 0 ; i < len(tilesResult.Val()) ; i++ {
			if i < 20 {
				var resultMovie RecommendationService.ContentTile
				err := proto.Unmarshal([]byte(fmt.Sprintf("%v", tilesResult.Val()[i])), &resultMovie)
				if err != nil {
					return err
				}
				err = stream.Send(&resultMovie)
				if err != nil {
					return err
				}
			}else {
				break
			}
		}
		return nil
	}else {
		err = h.PreProcessingContentbasedData(req)
		if err != nil {
			return err
		}
		result := h.RedisConnection.SMembers(fmt.Sprintf("user:%s:Contentbased", req.UserId))
		if result.Err() != nil {
			return result.Err()
		}
		tilesResult := h.RedisConnection.HMGet("cloudwalkerTiles", result.Val()...)
		for i := 0 ; i < len(tilesResult.Val()) ; i++ {
			if i < 20 {
				var resultMovie RecommendationService.ContentTile
				err := proto.Unmarshal([]byte(fmt.Sprintf("%v", tilesResult.Val()[i])), &resultMovie)
				if err != nil {
					return err
				}
				err = stream.Send(&resultMovie)
				if err != nil {
					return err
				}
			}else {
				break
			}
		}
		return nil
	}
}

func (h *RecommendationServiceHandler) PreProcessingContentbasedData(req *RecommendationService.GetRecommendationRequest) error  {

	log.Println("PreProcessingContentbasedData")
	// getting genre for user
	result := h.RedisConnection.SMembers(fmt.Sprintf("user:%s:genre", req.UserId))
	if result.Err() != nil {
		return result.Err()
	}
	var tempGenre []string
	for _,v := range result.Val() {
		log.Println("genre   ====>    ", v)
		tempGenre = append(tempGenre,fmt.Sprintf("genre:%s:items", formatString(v)))
	}
	genreResult := h.RedisConnection.SUnionStore(fmt.Sprintf("user:%s:genre:items", req.UserId),tempGenre...)
	log.Println("genre",genreResult.Val())
	// getting Languages for user
	result = h.RedisConnection.SMembers(fmt.Sprintf("user:%s:languages", req.UserId))
	if result.Err() != nil {
		return result.Err()
	}
	var tempLanguages []string
	for _,v := range result.Val() {
		log.Println("lang    ====>    ", v)
		tempLanguages = append(tempLanguages,fmt.Sprintf("languages:%s:items", formatString(v)) )
	}

	languageResult := h.RedisConnection.SUnionStore(fmt.Sprintf("user:%s:languages:items", req.UserId),tempLanguages...)
	log.Println("language ",languageResult.Val())
	// getting categories for user
	result = h.RedisConnection.SMembers(fmt.Sprintf("user:%s:categories", req.UserId))
	if result.Err() != nil {
		return result.Err()
	}
	var tempCategories []string
	for _,v := range result.Val() {
		log.Println("cateogires   ====>    ", v)
		tempCategories = append(tempCategories,fmt.Sprintf("categories:%s:items", formatString(v)) )
	}


	categoriesResult := h.RedisConnection.SUnionStore(fmt.Sprintf("user:%s:categories:items", req.UserId),tempCategories...)
	log.Println("categories ",categoriesResult.Val())

	// get tiles which content all its taste of user
	unionResultCount := h.RedisConnection.SUnionStore(fmt.Sprintf("user:%s:UnionStore", req.UserId),
		fmt.Sprintf("user:%s:genre:items", req.UserId),
		fmt.Sprintf("user:%s:languages:items", req.UserId),
		fmt.Sprintf("user:%s:categories:items", req.UserId)).Val()

	log.Println("Union Result count    ", unionResultCount)

	tileClickedByUser := h.RedisConnection.ZRevRange(fmt.Sprintf("user:%s:items", req.UserId), 0, -1)
	if len(tileClickedByUser.Val()) > 0 {
		h.RedisConnection.SAdd(fmt.Sprintf("user:%s:temp", req.UserId),tileClickedByUser.Val())
		resultOfDiff := h.RedisConnection.SDiffStore(fmt.Sprintf("user:%s:Contentbased", req.UserId),fmt.Sprintf("user:%s:UnionStore", req.UserId),fmt.Sprintf("user:%s:temp", req.UserId))
		if resultOfDiff.Err() != nil {
			log.Fatal(resultOfDiff.Err())
		}
		log.Println(resultOfDiff.Val())
	}else {
		interTileUser := h.RedisConnection.SMembers(fmt.Sprintf("user:%s:UnionStore", req.UserId))
		h.RedisConnection.SAdd(fmt.Sprintf("user:%s:Contentbased", req.UserId),interTileUser.Val())
	}

	//Deleting temp key
	h.RedisConnection.Del(fmt.Sprintf("user:%s:temp", req.UserId), fmt.Sprintf("user:%s:UnionStore", req.UserId))
	return nil
}

func(h *RecommendationServiceHandler) DiffContentBaseToTileClickedByUser(userId string){
	tileClickedByUser := h.RedisConnection.ZRevRange(fmt.Sprintf("user:%s:items", userId), 0, -1)
	h.RedisConnection.SAdd(fmt.Sprintf("user:%s:temp", userId),tileClickedByUser.Val())
	log.Println("Before ", h.RedisConnection.SCard(fmt.Sprintf("user:%s:Contentbased", userId)))
	h.RedisConnection.SDiffStore(fmt.Sprintf("user:%s:Contentbased", userId), fmt.Sprintf("user:%s:Contentbased", userId),fmt.Sprintf("user:%s:temp", userId))
	log.Println("After ", h.RedisConnection.SCard(fmt.Sprintf("user:%s:Contentbased", userId)))
	h.RedisConnection.Del(fmt.Sprintf("user:%s:temp", userId))
}

func (h *RecommendationServiceHandler) CheckInRedis(redisKey string) (bool, error) {
	intCmdResult := h.RedisConnection.Exists(redisKey)
	if intCmdResult.Val() == 1 {
		return  true , nil
	}else {
		return false , nil
	}
}

