package main

import (
	"context"
	"github.com/vallerion/pingpong-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func main() {
	con, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer con.Close()

	cli := proto.NewMultiplayerClient(con)

	ctx := context.Background()
	stream, err := cli.GameProcess(ctx)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	ch := make(chan error)

	go func(stream proto.Multiplayer_GameProcessClient) {
		for {
			err := stream.Send(&proto.Request{Action: &proto.Request_NewPlayer{&proto.NewPlayer{Id: "123"}}})
			if err != nil {
				log.Fatalf("failed to send message: %v", err)
			}
			time.Sleep(time.Second * 2)
		}
	}(stream)

	go func(stream proto.Multiplayer_GameProcessClient) {
		for {
			err := stream.Send(&proto.Request{Action: &proto.Request_Ping{&proto.Ping{StartTime: timestamppb.Now()}}})
			if err != nil {
				log.Fatalf("failed to send message: %v", err)
			}
			time.Sleep(time.Second * 1)
		}
	}(stream)

	go func(stream proto.Multiplayer_GameProcessClient) {
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Fatalf("failed to receive response: %v", err)
			}

			switch res.Action.(type) {
			case *proto.Response_Pong:
				pingTime := time.Since(res.GetPong().StartTime.AsTime()).Milliseconds()
				log.Printf("ping %d ms", pingTime)
			case *proto.Response_Players:
				log.Printf("received, leftPlayer.Id: %v, rightPlayer.Id: %v", res.GetPlayers().LeftPlayer.Id, res.GetPlayers().RightPlayer.Id)
			}
		}
	}(stream)

	<-ch
}
