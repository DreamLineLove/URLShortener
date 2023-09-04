package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "source.yml", "The name of the yaml source file")

	mux := http.NewServeMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandlerFn := MapHandler(pathsToUrls, mux)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	var yamlBytes []byte
	file.Read(yamlBytes)

	yamlHandlerFn, err := YAMLHandler(yamlBytes, mapHandlerFn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting the server on 8080...")
	http.ListenAndServe(":8080", yamlHandlerFn)
}
