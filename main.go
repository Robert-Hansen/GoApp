package main

import (
	"github.com/robert-hansen/goapp/routes"
	"github.com/robert-hansen/goapp/app"
	"github.com/robert-hansen/goapp/config"
	"log"
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