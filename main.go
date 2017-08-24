package main

import (
	"fmt"
	"github.com/robert-hansen/goapp/routes"
	"log"
	"github.com/robert-hansen/goapp/middleware"
)


func main() {
	fmt.Println("hejsdfsdsa123")
	router := routes.NewRouter()
	log.Println(router.HandleOPTIONS)
	hej := middleware.Hej{"hej"}
	log.Println(hej)
}