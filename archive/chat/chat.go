package intensive_go

import (
	"context"
	"log"
)

type Server struct{ UnimplementedChatServiceServer }

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Request from client: %s\n", message.Body)
	return &Message{
		Body: "Hello world from SayHello server",
	}, nil
}
