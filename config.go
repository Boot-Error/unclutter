package main

import (
	"bytes"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Buckets map[string][]string `yaml:"buckets"`
}

// read config.yaml
func LoadConfig(configPath string) Config {
	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	log.Printf("Reading from %s", configPath)

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)

	var config Config
	err = yaml.Unmarshal(buf.Bytes(), &config)

	if err != nil {
		log.Fatalf("YAML Error %v", err)
	}

	return config
}
