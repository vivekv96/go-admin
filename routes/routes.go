package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vivekv96/go-admin/handlers"
	"github.com/vivekv96/go-admin/middlewares"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", handlers.Register)
	app.Post("/api/login", handlers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/user", handlers.User)
	app.Post("/api/logout", handlers.Logout)
}
