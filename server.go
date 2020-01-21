package main


import (
	"context"
	"fmt"
	codecs "github.com/amsokol/mongo-go-driver-protobuf"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/transavro/RecommenderService/apihandler"
	pb "github.com/transavro/RecommenderService/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	atlasMongoHost          = "mongodb://nayan:tlwn722n@cluster0-shard-00-00-8aov2.mongodb.net:27017,cluster0-shard-00-01-8aov2.mongodb.net:27017,cluster0-shard-00-02-8aov2.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority"
	//developmentMongoHost = "mongodb://dev-uni.cloudwalker.tv:6592"
	developmentMongoHost = "mongodb://192.1.168.9:27017"
	schedularMongoHost   = "mongodb://192.168.1.143:27017"
	schedularRedisHost   = ":6379"
	grpc_port        = ":7765"
	rest_port		 = ":7766"
	srvCertFile = "cert/server.crt"
	srvKeyFile  = "cert/server.key"
)

// private type for Context keys
type contextKey int

const (
	clientIDKey contextKey = iota
)

var scheduleCollection, tileCollection *mongo.Collection
var tileRedis *redis.Client

// Multiple init() function
func init() {
	fmt.Println("Welcome to init() function")
	scheduleCollection = getMongoCollection("cloudwalker", "schedule", developmentMongoHost)
	tileCollection = getMongoCollection("cwtx2devel", "tiles", developmentMongoHost)
	tileRedis = getRedisClient(schedularRedisHost)
}

func credMatcher(headerName string) (mdName string, ok bool) {
	if headerName == "Login" || headerName == "Password" {
		return headerName, true
	}
	return "", false
}

// authenticateAgent check the client credentials
func authenticateClient(ctx context.Context, s *apihandler.RecommendationServiceHandler) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		clientLogin := strings.Join(md["login"], "")
		clientPassword := strings.Join(md["password"], "")
		if clientLogin != "nayan" {
			return "", fmt.Errorf("unknown user %s", clientLogin)
		}
		if clientPassword != "makasare" {
			return "", fmt.Errorf("bad password %s", clientPassword)
		}
		log.Printf("authenticated client: %s", clientLogin)
		return "42", nil
	}
	return "", fmt.Errorf("missing credentials")
}

//uthenticate client using jwt
func auth(ctx context.Context) error {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(
			codes.InvalidArgument,
			"missing context",
		)
	}

	authString, ok := meta["authorization"]
	if !ok {
		return status.Errorf(
			codes.Unauthenticated,
			"missing authorization",
		)
	}
	// validate token algo
	log.Println("found jwt token")
	jwtToken, err := jwt.Parse(
		authString[0],
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("bad signing method")
			}
			// additional validation goes here.
			return []byte("transavro"), nil
		},
	)

	if jwtToken.Valid {
		return nil
	}
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return status.Error(codes.Internal, "bad token")
}

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	s, ok := info.Server.(*apihandler.RecommendationServiceHandler)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("unable to cast the server"))
	}
	clientID, err := authenticateClient(ctx, s)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, clientIDKey, clientID)
	return handler(ctx, req)

	// auth using jwt

	//if err := auth(ctx); err != nil {
	//	return nil, err
	//}
	//log.Println("authorization OK")
	//return handler(ctx, req)
}

// streamAuthIntercept intercepts to validate authorization
func streamAuthIntercept(
	server interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {

	// auth using jwt
	if err := auth(stream.Context()); err != nil {
		return err
	}
	log.Println("authorization OK")
	return handler(server, stream)
}

func startGRPCServer(address, certFile, keyFile string) error {
	// create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// create a server instance

	redrecObj,err := apihandler.New(":6379")
	if err != nil {
		log.Println("nayan",err)
	}

	s := apihandler.RecommendationServiceHandler{
		RedisConnection:           getRedisClient(schedularRedisHost),
		CloudwalkerRecommendation: redrecObj,
		MongoCollection:           getMongoCollection( "test", "cwmovies", atlasMongoHost),
	}


	// Create the TLS credentials
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		return fmt.Errorf("could not load TLS keys: %s", err)
	}

	// Create an array of gRPC options with the credentials
	_ = []grpc.ServerOption{grpc.Creds(creds), grpc.UnaryInterceptor(unaryInterceptor), grpc.StreamInterceptor(streamAuthIntercept)}

	// create a gRPC server object
	//grpcServer := grpc.NewServer(opts...)

	// attach the Ping service to the server
	grpcServer := grpc.NewServer()                    // attach the Ping service to the server
	pb.RegisterRecommendationServiceServer(grpcServer, &s) // start the server
	log.Printf("starting HTTP/2 gRPC server on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}
	return nil
}

func startRESTServer(address, grpcAddress, certFile string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(credMatcher),)

	//creds, err := credentials.NewClientTLSFromFile(certFile, "")
	//if err != nil {
	//	return fmt.Errorf("could not load TLS certificate: %s", err)
	//}  // Setup the client gRPC options
	//
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(creds),}  // Register ping

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

	//grpcAddress := fmt.Sprintf("%s:%d", "cloudwalker.services.tv", 7775)
	//restAddress := fmt.Sprintf("%s:%d", "cloudwalker.services.tv", 7776)

	// fire the gRPC server in a goroutine
	go func() {
		err := startGRPCServer(grpc_port, srvCertFile, srvKeyFile)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	// fire the REST server in a goroutine
	go func() {
		err := startRESTServer(rest_port, grpc_port, srvCertFile)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	// infinite loop
	log.Printf("Entering infinite loop")
	select {}
}