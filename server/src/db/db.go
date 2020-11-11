package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DbConnect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=Niko1004 dbname=cafe-share sslmode=disable")

	// db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=Niko1004 dbname=cafe-share sslmode=disable")

	if err != nil {
		panic(err.Error())
	}
	return db
}
