package main

import (
	"fmt"
	"grpc/hello"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9099", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := hello.NewPingClient(conn)
	response, err := c.SayHello(context.Background(), &hello.PingMessage{Greeting: "You say goodbye!"})
	if err != nil {
		panic(err)
	}
	fmt.Println("Resposta:", response.Greeting)
}
