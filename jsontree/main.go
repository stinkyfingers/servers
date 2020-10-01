package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"path"
	"reflect"
	"strings"
)

// main
func main() {
	err := runServer()
	if err != nil {
		log.Fatal(err)
	}
}

// config files
var names = []string{
	"food.json",
	"sports.json",
}

var ErrNotFound = errors.New("config not found")

// preload config file
type FileManager struct {
	Files map[string]map[string]interface{} `json:"files"`
}

func NewFileManager() *FileManager {
	return &FileManager{
		Files: make(map[string]map[string]interface{}),
	}
}

func (f *FileManager) loadFile(filename string) error {
	var fileContents map[string]interface{}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&fileContents)
	if err != nil {
		return err
	}
	f.Files[strings.TrimSuffix(filename, path.Ext(filename))] = fileContents // TODO use map with name:file rather than trimmed name
	return nil
}

func (f *FileManager) loadFiles() error {
	for _, name := range names {
		err := f.loadFile(name)
		if err != nil {
			return err
		}
	}
	return nil
}

// server & handlers

func runServer() error {
	f := NewFileManager()
	err := f.loadFiles()
	if err != nil {
		return err
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", f.fileHandler)
	return http.ListenAndServe(":7000", mux)
}

func (f *FileManager) fileHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	config := paths[0]
	remaining := paths[1:]
	resp, err := f.traverseConfig(config, remaining)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	j, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(j)
}

// config getter
func (f *FileManager) traverseConfig(config string, paths []string) (interface{}, error) {
	output := struct {
		Data interface{} `json:"data"` // to handle []interface{} and interface{]}
	}{}
	file, ok := f.Files[config]
	if !ok {
		return nil, ErrNotFound
	}
	current := file
	for _, path := range paths {
		if _, ok = current[path]; !ok {
			return nil, ErrNotFound
		}
		if reflect.TypeOf(current[path]).Kind() != reflect.Map {
			output.Data = current[path]
			return output, nil
		}
		current, ok = current[path].(map[string]interface{})
		if !ok {
			return nil, ErrNotFound
		}
	}
	output.Data = current
	return output, nil
}
