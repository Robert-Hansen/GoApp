package app

import (
	"fmt"
	"log"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/robert-hansen/goapp/config"
	"github.com/julienschmidt/httprouter"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type App struct {
	Config 		config.Config
	Database	*gorm.DB
}

func New(cfg config.Config) *App {
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		cfg.MySQL.Username,
		cfg.MySQL.Password,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.DatabaseName,
		cfg.MySQL.Encoding)
	db, err := gorm.Open("mysql", dbConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return &App{cfg, db}
}

func (a *App) Run(r *httprouter.Router) {
	port := a.Config.Port
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("GOAPP is listening on port: %d\n", port)
	log.Fatal(http.ListenAndServe(addr, r))
}