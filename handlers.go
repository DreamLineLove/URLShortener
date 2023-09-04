package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if redirectUrl, exists := pathsToUrls[r.URL.Path]; exists {
			http.Redirect(w, r, redirectUrl, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte) ([]map[interface{}]interface{}, error) {
	var m []map[interface{}]interface{}
	err := yaml.Unmarshal(yml, &m)
	return m, err
}

func buildMap(data []map[interface{}]interface{}) map[string]string {
	m := make(map[string]string)
	for _, v := range data {
		for key, value := range v {
			keyStr := fmt.Sprintf("%v", key)
			valueStr := fmt.Sprintf("%v", value)
			m[keyStr] = valueStr
		}
	}
	return m
}
