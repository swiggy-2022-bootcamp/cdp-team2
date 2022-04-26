package ports

type IServer interface {
	Initialize()
	InitializeGrpcServer()
}
