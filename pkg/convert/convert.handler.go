package convert

import (
	"net/http"
	"strings"

	"github.com/Leonardo-Antonio/api_html-to-pdf/pkg/utils/res"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gofiber/fiber/v2"
)

type handler struct{}

func (h *handler) HtmlString(ctx *fiber.Ctx) error {
	body := new(body)
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(res.Err(err.Error(), err))
	}

	pdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(res.Err(err.Error(), err))
	}

	pdf.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(body.Content)))

	if err := pdf.Create(); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(res.Err(err.Error(), err))
	}

	if err := pdf.WriteFile("internal/static/template.pdf"); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(res.Err(err.Error(), err))
	}
	return ctx.Status(http.StatusOK).SendFile("internal/static/template.pdf")
}
