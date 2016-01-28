package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var running2 = true

func main() {
	log.Println("----------------")
	log.Println("starting up...")
	log.Println("----------------")

	go doSomething()

	// listen for signal(s)
	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel)

	// this service only acts on TERM, but more could be added here
	for ;running2; {
		s := <- sigChannel
		log.Printf("Signal Received: %+v", s)
		if s == syscall.SIGTERM {
			log.Println("shutting down...")
			running2 = false
		}
	}
}

func doSomething() {
	for ;running2; {
		time.Sleep(2 * time.Second)
		log.Println("Still here...")
	}
}
