package main

import (
	"net/http"
	"log"
	"github.com/robert-hansen/goapp/app"
	"github.com/robert-hansen/goapp/routes"
)


func main() {
	app := app.New()
	router := routes.NewRouter()
	app.Run(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}