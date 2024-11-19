package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-api-native/models"
)

var DB *gorm.DB

// ConnectDB connects to the database and sets up the database connection for the
// application. The function sets the database connection string using the
// environment variables set in the .env file. If the connection fails, the
// function panics with an appropriate error message. On successful connection,
// the function migrates the Author model to the database using the GORM
// AutoMigrate method. Finally, the function logs a success message to the
// console.
func ConnectDB() {
	connection := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Asia%vJakarta", ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_HOST, ENV.DB_PORT, ENV.DB_DATABASE, "%2F")

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database...")
	}

	db.AutoMigrate(&models.Author{})

	DB = db
	log.Println("Database connected...")
}
