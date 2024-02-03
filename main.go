package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	gormLogger "gorm.io/gorm/logger"
	"server.com/app/models"
	mdware "server.com/middleware"
	configs "server.com/pkg/config"
	"server.com/pkg/router"
	"server.com/platform/database"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Cannot load env")
	}
}

func autoMigrateDb() {
	db, err := database.OpenDBConnection()
	if err != nil {
		log.Fatal("Failed to connect to the DB\n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	db.Logger = gormLogger.Default.LogMode(gormLogger.Info)

	log.Println("Running Migrations")
	db.AutoMigrate(&models.User{})

	log.Println("Running Migrations end")

	if err != nil {
		log.Fatal("Failed to connect to the DB\n", err.Error())
		os.Exit(2)
	}
}

func main() {
	loadEnv()
	if os.Getenv("AUTO_MIGRATE") == "yes" {
		autoMigrateDb()
	}

	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	mdware.FiberMiddleware(app) // Register Fiber's middleware for app.
	// Router
	router.SetupRoutes(app)

	if err := app.Listen(":8001"); err != nil {
		log.Fatal(err)
	}
}
