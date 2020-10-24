package service

import (
	"context"
	"errors"
	"golang-grpc-examples/messages/messenger"
	"log"
)

type MessengerService struct {
	message *messenger.Message
}

func NewMessengerService() messenger.MessengerServiceServer {
	return &MessengerService{&messenger.Message{Body: "Hello from Server!"}}
}

func (s *MessengerService) WriteMessage(_ context.Context, req *messenger.Message) (*messenger.Empty, error) {

	if req.Body == "" {
		return nil, errors.New("Message.Body can't be empty")
	}

	log.Println("Received new message: ", req.Body)

	s.message.Body = req.Body

	return &messenger.Empty{}, nil
}

func (s *MessengerService) ReadMessage(_ context.Context, _ *messenger.Empty) (*messenger.Message, error) {
	return s.message, nil
}
