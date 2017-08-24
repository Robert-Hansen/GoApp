package main

import (
	"net/http"
	"log"
	"github.com/robert-hansen/GoApp/routes"
	"github.com/robert-hansen/GoApp/app"
	"github.com/robert-hansen/GoApp/config"
)


func main() {
	cfg, err := config.New("config/app.json")
	if err != nil {
		log.Fatal(err)
	}
	app := app.New(cfg)
	router := routes.NewRouter()
	app.Run(router)


	log.Fatal(http.ListenAndServe(":8080", router))
}