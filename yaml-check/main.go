package main

import (
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type EbExtension struct {
	Files    map[string]interface{} `yaml:"files"`
	Sources  map[string]interface{} `yaml:"sources"`
	Commands map[string]interface{} `yaml:"commands"`
}

var dir = flag.String("dir", "", "Path to directory with YAML files to run checker against")

func main() {
	flag.Parse()

	files, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(filepath.Join(*dir, file.Name()))
		if err != nil {
			log.Fatalf("error in file: %s : %v", file.Name(), err)
		}

		var e EbExtension
		if err := yaml.Unmarshal(data, &e); err != nil {
			log.Fatalf("error in file: %s : %v", file.Name(), err)
		}
	}
}
