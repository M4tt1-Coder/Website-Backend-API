package dblogger

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB               *gorm.DB
	connectionString = "root:MySqLt3sT25#@tcp(127.0.0.1:3306)/website?charset=utf8&parseTime=true&loc=Local"
) //!!!!change connection string when its finished

func Connect() {
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

func Init() {
	Connect()
}
