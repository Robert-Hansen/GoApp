package app

import (
	"fmt"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/robert-hansen/goapp/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/robert-hansen/goapp/database"
	"github.com/robert-hansen/goapp/database/migrations"
)

type Application struct {
	Config 		config.Config
	Database	*database.MySQLDB
}

func New(cfg config.Config) *Application {
	db, err := database.NewMYSQLDB(cfg.MySQL)
	if err != nil {
		log.Fatal(err)
	}
	migrations.RunUserMigration(db.DB)
	return &Application{cfg, db}
}

func (App *Application) Run(Router *httprouter.Router) {
	port := App.Config.Port
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("GOAPP is listening on port: %d\n", port)
	log.Fatal(http.ListenAndServe(addr, Router))
}