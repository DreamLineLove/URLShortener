package main

import (
	"net/http"

	"encoding/json"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if path, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, path, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yamldata []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathsToUrls []struct {
		Path string `yaml:"path"`
		URL  string `yaml:"url"`
	}
	if err := yaml.Unmarshal(yamldata, &pathsToUrls); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		for _, pathtourl := range pathsToUrls {
			if pathtourl.Path == r.URL.Path {
				http.Redirect(w, r, pathtourl.URL, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
}

func JSONHandler(jsondata []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathsToUrls []struct {
		Path string `json:"path"`
		URL  string `json:"url"`
	}
	if err := json.Unmarshal(jsondata, &pathsToUrls); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		for _, pathtourl := range pathsToUrls {
			if pathtourl.Path == r.URL.Path {
				http.Redirect(w, r, pathtourl.URL, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
}
