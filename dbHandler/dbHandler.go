package dbhandler

import (
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// The global database connection instance
	DB *gorm.DB
	// Connectionstring to the database
	connectionString = "root:MySqLt3sT25#@tcp(127.0.0.1:3306)/website?charset=utf8&parseTime=true&loc=Local"
) //!!!!change connection string when its finished

// Establishs the database connection
//
// Panics if the connection failed to be established.
func Connect() {
	//db, err := gorm.Open("mysql", connectionString)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}

// Returns the database connection
func GetDB() *gorm.DB {
	return DB
}
