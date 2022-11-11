package routes

import (
	"github.com/gofiber/fiber/v2"
	"api/ex/v2/controllers"
)

func SetUp(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}