package game

import (
	"context"

	"connectrpc.com/connect"
	gamev1 "github.com/imMatheus/syntax-city/server/gen/game/v1"
	"github.com/imMatheus/syntax-city/server/gen/game/v1/gamev1connect"
)

type gameServer struct{}

func NewGameServer() gamev1connect.GameServiceHandler {
	return &gameServer{}
}

func (g gameServer) GetGame(_ context.Context, req *connect.Request[gamev1.GetGameRequest]) (*connect.Response[gamev1.GetGameResponse], error) {
	return connect.NewResponse(&gamev1.GetGameResponse{
		GameBoard: &gamev1.GameBoard{
			Id: "1",
			Players: []*gamev1.Player{
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
	}), nil
}

// func StartGameServer(lis net.Listener) error {
// 	serverRegistrar := grpc.NewServer()
// 	service := &Server{}
// 	RegisterGameServer(serverRegistrar, service)

// 	err := serverRegistrar.Serve(lis)
// 	if err != nil {
// 		log.Fatalf("impossible to serve game server: %s", err)
// 	}

// 	return nil
// }

// 1. Really like the buf lint step
// 2. Dont really understand the v1 thing
// 3. Have a link to a full repo example with greet server
