package config

import (
	"fmt"
	"go-latihan/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func init() {
	InitDB()
	InitialMigration()
}

func InitDB() {

	config := Config{
		DB_Username: "asiyah",
		DB_Password: "asiyah354",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "middleware",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{}, &models.Book{}, &models.Blog{})
}
