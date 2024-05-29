package main

import (
    "strings"
	"flag"
	"fmt"
    "os"
	"net/http"
)

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The page you have requested does not exist!")
}

func main() {
    var path string
	flag.StringVar(&path, "path", "redirect.yml", "path to json or yaml file containing redirect URLs")

	flag.Parse()

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/me": "https://github.com/DreamLineLove",
        "/idol": "https://github.com/ThePrimeagen",
	}

	mapHandler := MapHandler(pathsToUrls, mux)

    if strings.HasSuffix(path, ".json") {
		jsonData, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		// fallback
		jsonHandler, err := JSONHandler([]byte(jsonData), mapHandler)
		if err != nil {
			panic(err)
		}

		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", jsonHandler)
    } else if strings.HasSuffix(path, ".yml") {
		yamlData, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		// fallback
		yamlHandler, err := YAMLHandler([]byte(yamlData), mapHandler)
		if err != nil {
			panic(err)
		}

		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", yamlHandler)
    } else {
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", mapHandler)
    }
}
