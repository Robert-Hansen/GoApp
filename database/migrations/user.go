package migrations

import (
	"log"
	"time"
	"math/rand"
	"github.com/jinzhu/gorm"
	"github.com/wawandco/fako"
	"github.com/robert-hansen/goapp/models"
)

func RunUserMigration(db *gorm.DB) {
	log.Println("Running User Migration")
	db.DropTableIfExists(&models.User{})
	db.AutoMigrate(&models.User{})

	for i := 0; i < 20; i++ {
		var user = &models.User{}
		fako.Fill(user)

		user.CreatedAt = randate(2000, 2017)
		user.Lastlogin = randate(2001, 2016)
		db.Create(user)
	}
}

func randate(from int, to int) time.Time {
	min := time.Date(from, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(to, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}