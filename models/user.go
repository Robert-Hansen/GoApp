package models

import (
	"time"
	"github.com/go-sql-driver/mysql"
)

/**
 * Users Table
 */
type User struct {
	ID			uint		`gorm:"primary_key"`
	Username	string		`gorm:"type:varchar(100);unique_index"fako:"user_name"`
	Email		string		`gorm:"type:varchar(100);unique_index"fako:"email_address"`
	Password	string		`gorm:"type:varchar(100);"fako:"simple_password"`
	Lastlogin	time.Time
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	mysql.NullTime	`sql:"index"`
}

/**
 * Name of database table
 */
func (User) TableName() string {
	return "users"
}
