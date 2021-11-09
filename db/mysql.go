package db

import (
	"fmt"
	"github.com/sachinsmc/ministry-of-programming-tasks/models"

	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	host := viper.GetString("db.host")
	dbname := viper.GetString("db.name")
	user := viper.GetString("db.username")
	password := viper.GetString("db.password")
	port := viper.GetString("db.port")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password,host, port, dbname)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database!")
	}

	fmt.Println("Connected to database.")

	DB.Debug()

	DB.Debug().AutoMigrate(&models.Product{}, &models.Admin{})

}

func GetDB() *gorm.DB {
	return DB
}
