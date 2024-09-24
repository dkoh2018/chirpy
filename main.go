package main

import (
	"log"
	"net/http"
)

func main() {
	port := "8080"
	filepathRoot := "./assets"

	serveMux := http.NewServeMux()

	// serving static files from /assets/
	serveMux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(filepathRoot))))

	// serve the index.html at the root url
	serveMux.Handle("/", http.FileServer(http.Dir(".")))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}