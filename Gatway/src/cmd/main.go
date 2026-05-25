package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/douglas/orbyte-gateway/src/routes"
)

func main() {

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
