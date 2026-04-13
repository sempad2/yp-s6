package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "HTTP: ", log.LstdFlags)
	serv := server.New(logger)

	logger.Println("Starting server on :8080")
	if err := serv.HTTPServer.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
