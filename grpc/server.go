package main

import (
	"net"

	"grpc/hello"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9099")
	if err != nil {
		panic(err)
	}
	// cria a instância do servidor
	s := hello.Server{}

	grpcServer := grpc.NewServer()

	hello.RegisterPingServer(grpcServer, &s)
	// inicia servidor
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
