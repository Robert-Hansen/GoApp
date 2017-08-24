package config

import (
	"os"
	"log"
	"encoding/json"
)

type MySQLConfig struct {
	Username     	string `json:"username"`
	Password 		string `json:"password"`
	DatabaseName    string `json:"database"`
	Encoding        string `json:"encoding"`
	Host        	string `json:"host"`
	Port        	string `json:"port"`
}

type Config struct {
	Env string 				`json:"env"`
	MySQL MySQLConfig 		`json:"mysql"`
	Port int 				`json:"port"`
}

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