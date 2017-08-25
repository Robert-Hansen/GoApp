package main

import (
	"log"
	"github.com/robert-hansen/goapp/app"
	"github.com/robert-hansen/goapp/routes"
	"github.com/robert-hansen/goapp/config"
)


func main() {
	cfg, err := config.New("config/app.json")
	if err != nil {
		log.Fatal(err)
	}
	app := app.New(cfg)
	router := routes.NewRouter()
	app.Run(router)
}