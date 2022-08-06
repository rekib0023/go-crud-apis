package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rekib0023/go-crud-apis/database"
	"github.com/rekib0023/go-crud-apis/routes"
)

func main() {
	database.Connect()

	app := fiber.New()
	api := app.Group("/api")

	routes.Setup(api)
	log.Fatal(app.Listen(":8000"))
}
