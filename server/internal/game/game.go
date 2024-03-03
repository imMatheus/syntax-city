package game

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	UnimplementedGameServer
}

func (g Server) GetGame(ctx context.Context, req *GetGameRequest) (*GetGameResponse, error) {
	return &GetGameResponse{
		GameBoard: &GameBoard{
			Id: "1",
			Players: []*Player{
				{
					Id:   "3",
					Name: "Player 1",
				},
				{
					Id:   "2",
					Name: "Player 2",
				},
			},
		},
	}, nil
}

func StartGameServer(lis net.Listener) error {
	serverRegistrar := grpc.NewServer()
	service := &Server{}
	RegisterGameServer(serverRegistrar, service)

	err := serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve game server: %s", err)
	}

	return nil
}
