package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	DEFAULT_BIND_ADDR = "0.0.0.0:8080"
	PORT              = kingpin.Flag("port", "Port to listen on (default: 8080)").Short('p').String()

	STATIC_DIR = kingpin.Arg("directory", "Directory to serve, defaults to current directory").ExistingDir()
)

func getDirectoryToServe() (string, error) {

	if STATIC_DIR != nil {
		return *STATIC_DIR, nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return wd, nil

}

func getBind() string {
	if PORT != nil {
		return ":" + *PORT
	}

	return DEFAULT_BIND_ADDR
}

func main() {
	kingpin.Parse()

	staticDir, err := getDirectoryToServe()
	if err != nil {
		log.Fatal("Unable to determine directory to serve")
	}

	fmt.Println("Listening on", getBind())
	log.Fatal(http.ListenAndServe(getBind(), http.FileServer(http.Dir(staticDir))))

}
