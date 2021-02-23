package hello

import (
	"fmt"

	"golang.org/x/net/context"
)

type Server struct {
	PingServer
}

func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	fmt.Println("Recebido:", in.Greeting)
	return &PingMessage{Greeting: "And I say hello!"}, nil
}
