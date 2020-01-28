package main

import (
	"context"
	"fmt"
	codecs "github.com/amsokol/mongo-go-driver-protobuf"
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pbAuth "github.com/transavro/AuthService/proto"
	"github.com/transavro/RecommenderService/apihandler"
	pb "github.com/transavro/RecommenderService/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	atlasMongoHost          = "mongodb://nayan:tlwn722n@cluster0-shard-00-00-8aov2.mongodb.net:27017,cluster0-shard-00-01-8aov2.mongodb.net:27017,cluster0-shard-00-02-8aov2.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority"
	developmentMongoHost = "mongodb://dev-uni.cloudwalker.tv:6592"
	//developmentMongoHost = "mongodb://192.1.168.9:27017"
	schedularMongoHost   = "mongodb://192.168.1.143:27017"
	schedularRedisHost   = ":6379"
	grpc_port        = ":7765"
	rest_port		 = ":7766"
)

// private type for Context keys
type contextKey int

const (
	clientIDKey contextKey = iota
)

var scheduleCollection, tileCollection *mongo.Collection
var tileRedis *redis.Client


//Temp struct
type EditorialTemp struct {
	ID struct {
		Oid string `json:"$oid"`
	} `json:"_id"`
	Title        string   `json:"title"`
	Portrait     []string `json:"portrait"`
	Poster       []string `json:"poster"`
	ContentID    string   `json:"contentId"`
	IsDetailPage bool     `json:"isDetailPage"`
	PackageName  string   `json:"packageName"`
	Target       []string `json:"target"`
}



// Multiple init() function
func init() {
	fmt.Println("Welcome to init() function")
	scheduleCollection = getMongoCollection("cloudwalker", "schedule", developmentMongoHost)
	tileCollection = getMongoCollection("cwtx2devel", "tiles", developmentMongoHost)
	tileRedis = getRedisClient(schedularRedisHost)

	go func() {

		//main init
		// creating pipes for mongo aggregation
		pipeline := mongo.Pipeline{}

		//Adding stage 2
		pipeline = append(pipeline, bson.D{{"$project", bson.D{
			{"title", "$metadata.title"},
			{"portrait", "$posters.portrait",},
			{"poster", "$posters.landscape"},
			{"contentId", "$ref_id"},
			{"isDetailPage", "$content.detailPage"},
			{"packageName", "$content.package"},
			{"target", "$content.target"},}}})

		// creating aggregation query
		tileCur, err := tileCollection.Aggregate(context.Background(), pipeline)
		if err != nil {
			log.Println(err.Error())
		}
		for tileCur.Next(context.Background()){

			var temp EditorialTemp
			var contentTile pb.ContentTile
			err = tileCur.Decode(&temp)
			if err != nil {
				log.Println(err.Error())
			}
			contentTile.ContentId = temp.ContentID
			contentTile.IsDetailPage = temp.IsDetailPage
			contentTile.PackageName = temp.PackageName
			if len(temp.Poster) > 0 {
				contentTile.Poster = temp.Poster[0]
			}
			if len(temp.Portrait) > 0 {
				contentTile.Portrait = temp.Portrait[0]
			}
			contentTile.Target = temp.Target
			contentTile.Title = temp.Title
			contentTile.TileType = pb.TileType_ImageTile

			if contentTile.XXX_Size() > 0 {
				contentByte, err := proto.Marshal(&contentTile)
				if err != nil {
					log.Println(err.Error())
				}
					if err = tileRedis.HSet("cloudwalkerTiles",contentTile.ContentId, contentByte).Err(); err != nil {
					log.Println(err.Error())
				}
			}
		}

	}()
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("unaryInterceptor")
	err := checkingJWTToken(ctx)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func checkingJWTToken(ctx context.Context) error{
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.NotFound, fmt.Sprintf("no auth meta-data found in request" ))
	}

	token := meta["token"]

	if len(token) == 0 {
		return  status.Error(codes.NotFound, fmt.Sprintf("Token not found" ))
	}

	// calling auth service
	conn, err := grpc.Dial(":7757", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Auth here
	authClient := pbAuth.NewAuthServiceClient(conn)
	_, err = authClient.ValidateToken(context.Background(), &pbAuth.Token{
		Token: token[0],
	})
	if err != nil {
		return  status.Error(codes.NotFound, fmt.Sprintf("Invalid token:  %s ", err ))
	}else {
		return nil
	}
}

// streamAuthIntercept intercepts to validate authorization
func streamIntercept(server interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler, ) error {
	err := checkingJWTToken(stream.Context())
	if err != nil {
		return err
	}
	return handler(server, stream)
}

func startGRPCServer(address string) error {
	// create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// create a server instance
	redrecObj,err := apihandler.New(":6379")
	if err != nil {
		return err
	}

	s := apihandler.RecommendationServiceHandler{
		RedisConnection:           getRedisClient(schedularRedisHost),
		CloudwalkerRecommendation: redrecObj,
		MongoCollection:           tileCollection,
	}

	serverOptions := []grpc.ServerOption{grpc.UnaryInterceptor(unaryInterceptor), grpc.StreamInterceptor(streamIntercept)}

	// attach the Ping service to the server
	grpcServer := grpc.NewServer(serverOptions...)
	// attach the Ping service to the server
	pb.RegisterRecommendationServiceServer(grpcServer, &s) // start the server
	log.Printf("starting HTTP/2 gRPC server on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}
	return nil
}

func startRESTServer(address, grpcAddress string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(runtime.DefaultHeaderMatcher))

	opts := []grpc.DialOption{grpc.WithInsecure()} // Register ping

	err := pb.RegisterRecommendationServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("could not register service Ping: %s", err)
	}

	log.Printf("starting HTTP/1.1 REST server on %s", address)
	http.ListenAndServe(address, mux)
	return nil
}

func getMongoCollection(dbName, collectionName, mongoHost string) *mongo.Collection {
	// Register custom codecs for protobuf Timestamp and wrapper types
	reg := codecs.Register(bson.NewRegistryBuilder()).Build()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoHost), options.Client().SetRegistry(reg))
	if err != nil {
		log.Fatal(err)
	}
	return mongoClient.Database(dbName).Collection(collectionName)
}

func getRedisClient(redisHost string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to redis %v", err)
	}
	return client
}

func main() {

	// fire the gRPC server in a goroutine
	go func() {
		err := startGRPCServer(grpc_port)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	// fire the REST server in a goroutine
	go func() {
		err := startRESTServer(rest_port, grpc_port)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	// infinite loop
	log.Printf("Entering infinite loop")
	select {}
}