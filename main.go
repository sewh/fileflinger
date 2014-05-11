package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

func defineFlags() (string, string) {
	var port, path string
	flag.StringVar(&port, "port", "6606", "The port for the server to listen on.")
	flag.StringVar(&path, "path", ".", "The path to the directory to serve.")
	path, err := filepath.Abs(path)

	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()

	return port, path
}

func main() {
	port, path := defineFlags()

	http.Handle("/", http.FileServer(http.Dir(path)))
	log.Printf("fileflinger starting on 0.0.0.0:%s for directory %s\n", port, path)
	log.Fatal(http.ListenAndServe("0.0.0.0:" + port, nil))

}
