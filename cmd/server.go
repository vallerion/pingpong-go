package main

import (
	"errors"
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
	clients map[uuid.UUID]proto.Multiplayer_GameProcessServer
	proto.UnimplementedMultiplayerServer
}

func (m *MultiplayerServer) GameProcess(stream proto.Multiplayer_GameProcessServer) error {
	currentClientId := uuid.New()
	m.clients[currentClientId] = stream
	fmt.Println(m.clients)

	for {
		req, err := stream.Recv()
		if err != nil {
			//delete(m.clients, currentClientId)
			return fmt.Errorf("failed to reveive: %v", err)
		}
		switch req.Action.(type) {
		case *proto.Request_Ping:
			err = stream.Send(&proto.Response{Action: &proto.Response_Pong{Pong: &proto.Pong{StartTime: req.GetPing().StartTime}}})
			if err != nil {
				log.Printf("failed to send response: %v", err)
			}
		case *proto.Request_NewPlayer:
			if m.game.LeftPlayer != nil && m.game.RightPlayer != nil {
				return errors.New("failed to connect new player")
			}
			if _, valid := uuid.Parse(req.GetNewPlayer().Id); valid != nil {
				return fmt.Errorf("received invalid player id: %v", valid)
			}

			if m.game.LeftPlayer == nil {
				m.game.LeftPlayer = &server.Player{
					Id: uuid.MustParse(req.GetNewPlayer().Id),
					X:  req.GetNewPlayer().X,
					Y:  req.GetNewPlayer().Y,
				}
			} else if m.game.RightPlayer == nil {
				m.game.RightPlayer = &server.Player{
					Id: uuid.MustParse(req.GetNewPlayer().Id),
					X:  req.GetNewPlayer().X,
					Y:  req.GetNewPlayer().Y,
				}
			}

			// broadcast other players
			for k, v := range m.clients {
				if k != currentClientId {
					err = v.Send(&proto.Response{Action: &proto.Response_Players{Players: &proto.Players{
						LeftPlayer:  &proto.Player{Id: m.game.LeftPlayer.Id.String(), X: m.game.LeftPlayer.X, Y: m.game.LeftPlayer.Y},
						RightPlayer: &proto.Player{Id: m.game.RightPlayer.Id.String(), X: m.game.RightPlayer.X, Y: m.game.RightPlayer.Y},
					}}})
					if err != nil {
						log.Printf("failed to send response: %v", err)
					}
				}
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

	game := &server.Game{Status: server.PENDING}
	//game := &server.Game{Rooms: make(map[uuid.UUID]*server.Room, 0)}

	s := grpc.NewServer()
	ms := &MultiplayerServer{game: game, clients: make(map[uuid.UUID]proto.Multiplayer_GameProcessServer)}
	proto.RegisterMultiplayerServer(s, ms)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
