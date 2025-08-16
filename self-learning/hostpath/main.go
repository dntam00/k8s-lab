package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	podName := os.Getenv("POD_NAME")
	log.Printf("POD_NAME env: %s", podName)

	// Ensure /data directory exists
	if err := os.MkdirAll("/data", 0755); err != nil {
		log.Fatalf("failed to create /data directory: %v", err)
	}

	// Open log file for writing
	logFile, err := os.OpenFile("/data/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Set log output to file
	log.SetOutput(logFile)

	log.Print("Starting the server test...")

	log.Fatal("error while starting the server")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigs
	log.Printf("Received signal: %v, shutting down.", sig)
}
