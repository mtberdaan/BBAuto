package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func eventLoop(eventChan chan string) {
	for {
		// generate an event
		event := time.Now().Format("2006-01-02 15:04:05")

		// sent the event to main loop
		eventChan <- event

		// simulate some work
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// channel to receive termination signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// channel for events
	eventChan := make(chan string)

	// start event loop in goroutine
	go eventLoop(eventChan)

	// Main event loop
	for {
		select {
		case event := <-eventChan:
			// process event
			log.Println("processing event:", event)

		case <-signalChan:
			// termination signal received
			log.Println("termination signal received. exiting...")
			return
		}
	}
}
