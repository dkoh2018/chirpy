package main

import (
	"log"
	"net/http"
)

func main() {
	port := "8080"
	filepathRoot := "./"
	
	serveMux := http.NewServeMux()
	serveMux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
