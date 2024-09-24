package main

import (
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.Handle("/", http.FileServer(http.Dir("./")))

	server := &http.Server{
		Handler: serveMux,
		Addr:    ":8080",
	}

	log.Println("Starting server on", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
