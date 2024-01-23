package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/kitpk/go-unittest106/adapters"
	"github.com/kitpk/go-unittest106/core"
)

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
			SlowThreshold: time.Second, // Slow gSQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	// Initialize the database connection
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect database")
	}

	// Migrate the schema
	if err := db.AutoMigrate(&core.Order{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}

func setup() *fiber.App {
	app := fiber.New()

	db := InitializeDB()

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderServicer(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)
	app.Post("/order", orderHandler.CreateOrder)

	return app
}

func main() {
	app := setup()
	app.Listen(":8080")
}
