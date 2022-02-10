package main

import (
	"log"

	"github.com/Leonardo-Antonio/api_html-to-pdf/pkg/convert"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	convert.Router(app)

	log.Fatalln(app.Listen(":8000"))
}
