package app

import (
	"log"
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

type App struct {

}

func New() *App {
	return &App{}
}

func (a *App) Run(r *httprouter.Router)  {
	port := 8080
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("APP is listening on port: %d\n", port)
	log.Fatal(http.ListenAndServe(addr, r))
}
