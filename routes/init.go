package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/csrf"
)

func InitializeRoutes(app *fiber.App) {

	app.Get("/*", func(c fiber.Ctx) error {

		csrf := csrf.TokenFromContext(c)

		return c.Render("app", fiber.Map{
			"title": "Home",
			"csrf":  csrf,
		})
	})

}
