package main

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

func main() {
	yml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	var m []map[interface{}]interface{}
	err := yaml.Unmarshal([]byte(yml), &m)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(m)
	}

	fmt.Println("Starting the server on 8080...")
	http.ListenAndServe(":8080", nil)
}
