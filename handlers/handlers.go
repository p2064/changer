package handlers

import (
	"errors"
	"log"
	"time"

	"github.com/p2064/changer/proto"
	"github.com/p2064/pkg/db"
	kafka "github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) ChangeEvent(ctx context.Context, in *proto.ChangeEventRequest) (
	*proto.ChangeEventResponse,
	error,
) {
	log.Printf("Receive message body from client: %s", in.String())
	data := db.Event{
		ID:         in.Id,
		Place:      in.Place,
		EventTime:  in.Time,
		MaxPlayers: in.MaxPlayers,
	}
	var updateEvent db.Event
	res := db.DB.Model(&updateEvent).Where("id = ?", in.Id).Updates(&data)
	if res.Error != nil {
		return &proto.ChangeEventResponse{Status: 400, Error: errors.New("Event not changed").Error()}, errors.New("Event not changed")
	}
	topic := "notify"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, partition)
	if err != nil {
		log.Printf("failed to dial leader: %v", err)
	} else {
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		_, err = conn.WriteMessages(
			kafka.Message{Value: []byte(in.String())},
		)
		if err != nil {
			log.Printf("failed to write messages: %v", err)
		}
		if err := conn.Close(); err != nil {
			log.Printf("failed to close writer: %v", err)
		}
	}
	return &proto.ChangeEventResponse{Status: 200, Error: "No error"}, nil
}
