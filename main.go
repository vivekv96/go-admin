package main

import (
	"log"

	"github.com/vivekv96/go-admin/database"

	"github.com/vivekv96/go-admin/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := database.ConnectToMySQL(&database.MySQLConfig{
		Host:     "localhost",
		Username: "root",
		Password: "root123",
		Port:     3306,
		DBName:   "admin",
	}); err != nil {
		log.Fatalln("Could not connect to database: ", err)
	}

	app := fiber.New()

	routes.Setup(app)

	if err := app.Listen(":8000"); err != nil {
		log.Panicln(err)
	}
}
