package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/vallerion/pingpong-go/proto"
	"github.com/vallerion/pingpong-go/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

type MultiplayerServer struct {
	game    *server.Game
	clients map[uuid.UUID]*proto.Multiplayer_GameProcessServer
	proto.UnimplementedMultiplayerServer
}

func (m *MultiplayerServer) GameProcess(stream proto.Multiplayer_GameProcessServer) error {
	id := uuid.New()
	m.clients[id] = &stream
	fmt.Println(m.clients)

	for {
		req, err := stream.Recv()
		if err != nil {
			log.Fatalf("failed to reveive: %v", err)
		}
		switch req.Action.(type) {
		case *proto.Request_Ping:
			err = stream.Send(&proto.Response{Action: &proto.Response_Pong{Pong: &proto.Pong{StartTime: req.GetPing().StartTime}}})
			if err != nil {
				log.Fatalf("failed to send response: %v", err)
			}
		}
	}

	return nil
}

func main() {
	port := flag.Int("port", 8888, "The port to listen on.")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %d", *port)

	game := &server.Game{Rooms: make(map[uuid.UUID]*server.Room, 0)}

	s := grpc.NewServer()
	ms := &MultiplayerServer{game: game, clients: make(map[uuid.UUID]*proto.Multiplayer_GameProcessServer)}
	proto.RegisterMultiplayerServer(s, ms)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
