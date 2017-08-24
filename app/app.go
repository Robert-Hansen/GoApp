package app

import (
	"log"
	"net/http"
	"fmt"
	"github.com/robert-hansen/GoApp/config"
	"github.com/julienschmidt/httprouter"
)

type App struct {
	Config 		config.Config
}

func New(cfg config.Config) *App {

}

func (a *App) Run(r *httprouter.Router)  {
	port := a.Config.Port
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("APP is listening on port: %d\n", port)
	log.Fatal(http.ListenAndServe(addr, r))
}

func (a *App) IsProd() bool {
	return a.Config.Env == "prod"
}