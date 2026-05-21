package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/csrf"
)

func InitializeRoutes(app *fiber.App) {

	app.Get("/", func(c fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Home",
		})
	})
	app.Post("/auth/login", func(c fiber.Ctx) {
		fmt.Println("OKE!")
	})
	app.Get("/auth/login", func(c fiber.Ctx) error {
		token := csrf.TokenFromContext(c)
		return c.Render("auth/login", fiber.Map{
			"title": "Login",
			"csrf":  token,
		})
	})
}
