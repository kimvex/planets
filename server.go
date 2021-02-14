package main

import (
	"planets/db"
	"planets/routes"

	"github.com/gofiber/fiber/v2"
)

var (
	apiRoute fiber.Router
	database *db.DB
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 200 * 1024 * 1024,
	})

	database = db.NewDatabase()
	database.CloseConnection()

	apiRoute = app.Group("/api")

	routes.API(apiRoute)

	app.Listen(":4004")
}
