package main

import (
	"bytes"
	_ "embed"
	"log"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var defaultConfigData string

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

func SetupConfigLocally() string {
	configPath := configdir.LocalConfig("unclutter")
	err := configdir.MakePath(configPath)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	configFile := filepath.Join(configPath, "config.yaml")

	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		fh, err := os.Create(configFile)
		if err != nil {
			log.Fatalf("Error %v", err)
		}
		defer fh.Close()

		_, err = fh.WriteString(defaultConfigData)
		if err != nil {
			log.Fatalf("Error %v", err)
		}

		fh.Sync()
	}

	return configFile
}
