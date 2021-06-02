package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vivekv96/go-admin/handlers"
)

func Setup(app *fiber.App) {
	app.Get("/", handlers.Hello)
}
