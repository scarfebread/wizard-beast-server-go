package main

import (
	"github.com/scarfebread/wizard-beast-server-go/internal/game"
	"github.com/scarfebread/wizard-beast-server-go/internal/player"
	"github.com/scarfebread/wizard-beast-server-go/internal/request"
	"github.com/scarfebread/wizard-beast-server-go/internal/udp"
	"log"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(2) // TODO how does this work with embedded goroutines?

	playerRepository := player.NewRepository()
	server := udp.Server{
		Processor: request.Processor{PlayerRepository: playerRepository},
	}
	engine := game.Engine{
		PlayerRepository: playerRepository,
	}

	go startServer(wg, server)
	go startEngine(wg, engine)

	wg.Wait()
}

func startServer(wg *sync.WaitGroup, server udp.Server) {
	defer wg.Done()

	err := server.Start()

	if err != nil {
		log.Fatal(err)
	}
}

func startEngine(wg *sync.WaitGroup, engine game.Engine) { // TODO this is duplicated but might make sense as I add more
	defer wg.Done()

	err := engine.Start()

	if err != nil {
		log.Fatal(err)
	}
}
