package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandlerFn := MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	yamlHandlerFn, err := YAMLHandler([]byte(yaml), mapHandlerFn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting the server on 8080...")
	http.ListenAndServe(":8080", yamlHandlerFn)
}
