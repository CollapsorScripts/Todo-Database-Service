package Service

import "google.golang.org/grpc"

type serverAPI struct {
	UnimplementedServiceServer
}

func Register(gRPC *grpc.Server) {
	RegisterServiceServer(gRPC, &serverAPI{})
}
