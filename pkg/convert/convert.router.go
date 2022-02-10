package convert

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {
	handler := &handler{}

	group := app.Group("/api/v1/convert")
	group.Post("/html-to-pdf", handler.HtmlString)
}
