package config

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// func BuildDBConfig() (dbconfig string) {

// 	dbConfig := "host=localhost user=root password=root dbname=db port=5432 sslmode=disable"

// 	return dbConfig
// }
