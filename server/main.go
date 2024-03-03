package main

import (
	"log"
	"net"

	"github.com/imMatheus/syntax-city/server/internal/game"
)

const PORT = ":8080"

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	log.Default().Printf("server is listening on port %s", PORT)

	err = game.StartGameServer(lis)

	if err != nil {
		log.Fatalf("cannot start game server: %s", err)
	}
}
