package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vivekv96/go-admin/database"
	"github.com/vivekv96/go-admin/routes"
)

func main() {
	if err := database.ConnectToMySQL(&database.MySQLConfig{
		Host:     "localhost",
		Username: "root",
		Password: "root123",
		Port:     3306,
		DBName:   "mysql",
	}); err != nil {
		log.Fatalln("Could not connect to database: ", err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	if err := app.Listen(":8000"); err != nil {
		log.Panicln(err)
	}
}
