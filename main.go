package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

//go:embed web/*
var embeddedFiles embed.FS

func main() {
	webFS, err := fs.Sub(embeddedFiles, "web")
	if err != nil {
		log.Fatalf("Failed to create subdirectory: %v", err)
	}

	mux := http.NewServeMux()
	files := http.FileServer(http.FS(webFS))

	// Handle static files
	mux.Handle("/", files)

	fmt.Println("Starting Server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
