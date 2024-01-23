package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"unique"`
	Age      int
}

func InitializeDB() *gorm.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "myuser"
		password = "mypassword"
		dbname   = "mydatabase"
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&User{})
	return db
}

func AddUser(db *gorm.DB, fullname string, email string, age int) error {
	user := User{Fullname: fullname, Email: email, Age: age}

	var count int64
	db.Model(&User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		return errors.New("email already exists")
	}

	result := db.Create(&user)
	return result.Error
}

func main() {
	db := InitializeDB()
	err := AddUser(db, "John Doe", "john.doe@example.com", 30)
	if err != nil {
		fmt.Println("err DB", err)
	}

}
