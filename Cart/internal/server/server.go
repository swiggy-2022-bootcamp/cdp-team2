package server

import (
	"net"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/config"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/server/grpchandlers"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/services"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/protos"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	Configuration *config.WebServerConfig
	Router        *Router
}

// The following handler will handle all the CORS related stuff.
func corsHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("inside cors handler")
		// origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			// To the tell browser on pre-flight request("OPTIONS") that transfer of cookie
			// is possible because "Access-Control-Allow-Credentials" is set to true.
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization, access-control-allow-origin")
			return
		}

		// The following header is set on all the request to prevent the following error:
		// Cross-Origin Request Blocked: The Same Origin Policy disallows reading the
		// remote resource at <server url>.
		// (Reason: expected ‘true’ in CORS header ‘Access-Control-Allow-Credentials’).
		// This header is more important when we need to send and receive "jwtToken"
		// cookie and "withCredentials" is set to true in request.
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		h.ServeHTTP(w, r)
	})
}

// NewServer creates the new server and sets the server configurations.
func NewServer(config *config.WebServerConfig) *Server {
	router := NewRouter()
	server := &Server{
		Configuration: config,
		Router:        router,
	}
	return server
}

// RunServer initializes the server along with all the microservice dependencies.
// It starts the server and returns nil as error if server starts successfully otherwise
// returns the error.
func RunServer() error {
	webServerConfig, err := config.FromEnv()
	if err != nil {
		return err
	}

	routerConfigs := util.RouterConfig{
		WebServerConfig: webServerConfig,
	}

	dynamoClient := createDynamoDbSession()
	dynamodao := dao.InitDynamoDAO(dynamoClient, webServerConfig)

	// Initialize services
	services.InitHealthCheckService(&routerConfigs)
	services.InitAddCartItemService(&routerConfigs, dynamodao)
	services.InitGetCartService(&routerConfigs, dynamodao)
	services.InitUpdateCartItemService(&routerConfigs, dynamodao)
	services.InitDeleteCartItemService(&routerConfigs, dynamodao)

	server := NewServer(webServerConfig)
	server.Router.InitializeRouter(&routerConfigs)

	go startGrpcServer()

	log.Info("Server starting on PORT: ", webServerConfig.Port)
	err = http.ListenAndServe(":"+webServerConfig.Port, corsHandler(*server.Router))
	if err != nil {
		return err
	}

	return nil
}

func startGrpcServer() {
	log.Info("Staring gRPC server on PORT: 9001")
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	protos.RegisterCartServiceServer(grpcServer, &grpchandlers.Server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func createDynamoDbSession() *dynamodb.DynamoDB {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	dbSession := session.New(&aws.Config{
		Region: aws.String("us-west-2"),
	})

	return dynamodb.New(dbSession)
}
