package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var running = true

func main() {
	log.Println("----------------")
	log.Println("starting up...")
	log.Println("----------------")

	go doStuff()

	// listen for exactly ONE signal
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGTERM)
	s := <- sigChannel
	log.Printf("Signal Received: %+v", s)
	running = false
	log.Println("shutting down...")
}

func doStuff() {
	for ;running; {
		time.Sleep(2 * time.Second)
		log.Println("Still here...")
	}
}