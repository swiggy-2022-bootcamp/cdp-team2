package server

import (
	"net"
	"net/http"

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

// NewServer creates the new server and sets the server configurations.
func NewServer(config *config.WebServerConfig) *Server {
	server := &Server{
		Configuration: config,
		Router:        NewRouter(),
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
	services.InithHealthCheckService(&routerConfigs)
	services.InitAddCartItemService(&routerConfigs, dynamodao)
	services.InitGetCartService(&routerConfigs, dynamodao)
	services.InitUpdateCartItemService(&routerConfigs, dynamodao)
	services.InitDeleteCartItemService(&routerConfigs, dynamodao)

	server := NewServer(webServerConfig)
	server.Router.InitializeRouter(&routerConfigs)

	go startGrpcServer()

	log.Info("Server starting on PORT: ", webServerConfig.Port)
	err = http.ListenAndServe(":"+webServerConfig.Port, *server.Router)
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
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	return svc
}
