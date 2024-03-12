package main

import (
	"log"
	"sync"
	"wizard-beast-server-go/engine"
	"wizard-beast-server-go/entity"
	"wizard-beast-server-go/server"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(2)

	playerRepository := entity.CreatePlayerRepository()
	server := server.Server{ // TODO server
		PlayerRepository: playerRepository,
	}

	go startServer(wg, server)
	go startEngine(wg)

	wg.Wait()
}

func startServer(wg *sync.WaitGroup, server server.Server) {
	defer wg.Done()

	err := server.Start()

	if err != nil {
		log.Fatal(err)
	}
}

func startEngine(wg *sync.WaitGroup) { // TODO this is duplicated but might make sense as I add more
	defer wg.Done()

	err := engine.Start()

	if err != nil {
		log.Fatal(err)
	}
}
