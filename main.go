package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "source.yml", `The name of the source file. File extension must be either ".yml" or ".json".`)

	flag.Parse()

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
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	sourceBytes := make([]byte, fileInfo.Size())
	_, err = file.Read(sourceBytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting the server on 8080...")
	switch strings.HasSuffix(filename, ".json") {
	case true:
		jsonHandlerFn, err := JSONHandler(sourceBytes, mapHandlerFn)
		if err != nil {
			log.Fatal(err)
		}
		http.ListenAndServe(":8080", jsonHandlerFn)
	default:
		yamlHandlerFn, err := YAMLHandler(sourceBytes, mapHandlerFn)
		if err != nil {
			log.Fatal(err)
		}
		http.ListenAndServe(":8080", yamlHandlerFn)
	}
}
