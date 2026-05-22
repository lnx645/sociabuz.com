package routes

import (
	"fmt"
	"html/template"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/csrf"
)

func InitializeRoutes(app *fiber.App) {
	app.Get("/*", func(c fiber.Ctx) error {
		figures := figure.NewFigure("Welcome To GoBuzz", "puffy", true)
		csrf := csrf.TokenFromContext(c)
		isDev := os.Getenv("SERVER_ENV") == "development"
		return c.Render("app", fiber.Map{
			"title":     "Home",
			"csrf":      csrf,
			"isDev":     isDev,
			"print_art": template.HTML(fmt.Sprintf("\n%s\n", figures.String())),
		})
	})

}
