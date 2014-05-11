// fileflinger:main.go
// Utility for serving a directory over HTTP
// Licensed under the three clause BSD license (see LICENSE):

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func defineFlags() (bool, string, string) {
	var port, path string
	var help bool

	flag.BoolVar(&help, "help", false, "Display fileflinger usage.")
	flag.StringVar(&port, "port", "6606", "The port for the server to listen on.")
	flag.StringVar(&path, "path", ".", "The path to the directory to serve.")
	path, err := filepath.Abs(path)

	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()

	return help, port, path
}

func main() {
	help, port, path := defineFlags()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	http.Handle("/", http.FileServer(http.Dir(path)))
	log.Printf("fileflinger starting on 0.0.0.0:%s for directory %s\n", port, path)
	log.Fatal(http.ListenAndServe("0.0.0.0:" + port, nil))

}
