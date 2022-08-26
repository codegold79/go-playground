package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kr/pretty"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Debug      bool                     `yaml:"debug"`
	UpdateFreq int                      `yaml:"updateFreq"`
	Faves      []map[string]interface{} `yaml:"faves"`
	Tools      map[string]interface{}   `yaml:"tools"`
}

func main() {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Printf("read config.yaml file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Printf("unmarshal config.yaml file: %v", err)
	}

	// Use the pretty package to make the Config struct easier to read.
	fmt.Printf("parsed config: %# v", pretty.Formatter(config))
}
