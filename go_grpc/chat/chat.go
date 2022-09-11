package chat

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct{}

// func (s *Server) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
// 	log.Println("Hello World: %s", in.Body)
// 	return &Message{Body: "Hello from server!"}, nil
// }

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Hello World: %s", message.Body)
	return &Message{Body: "Hello from server!"}, nil
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {
	log.Println("I dont know what this does")
}
