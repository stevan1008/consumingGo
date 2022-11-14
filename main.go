package main

import (
	"github.com/gofiber/fiber/v2"
	"api/ex/v2/database"
	"api/ex/v2/routes"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.SetUp(app)	
	app.Listen(":3100")
}