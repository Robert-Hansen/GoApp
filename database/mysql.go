package database

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/robert-hansen/goapp/config"
)

type MySQLDB struct {
	*gorm.DB
}

/**
 * NewMySQLDB creates and migrates the database
 */
func NewMYSQLDB(cfg config.MySQLConfig) (*MySQLDB, error) {
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
		cfg.Encoding)
	db, err := gorm.Open("mysql", dbConnection)
	if err != nil {
		panic("failed to connect database")
	}
	if err = db.DB().Ping(); err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)

	return &MySQLDB{db}, nil
}