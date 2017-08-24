package config

import (
	"os"
	"encoding/json"
	"log"
)

type Config struct {
	Env string 				`json:"env"`
	Port int 				`json:"port"`
}

// New creates a new config by reading a json file that matches the types above
func New(path string) (Config, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg, nil
}