package main

import (
	"log"
	"sync"
	"wizard-beast-server-go/engine"
	"wizard-beast-server-go/server"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go startServer(wg)
	go startEngine(wg)

	wg.Wait()
}

func startServer(wg *sync.WaitGroup) {
	defer wg.Done()

	err := server.Start()

	if err != nil {
		log.Fatal(err)
	}
}

func startEngine(wg *sync.WaitGroup) {
	defer wg.Done()

	err := engine.Start()

	if err != nil {
		log.Fatal(err)
	}
}
