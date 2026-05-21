package routes

import "github.com/gofiber/fiber/v3"

func InitializeRoutes(app *fiber.App) {

	app.Get("/", func(c fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Home",
		})
	})
	app.Get("/auth/login", func(c fiber.Ctx) error {
		return c.Render("auth/login", fiber.Map{
			"title": "Login",
		})
	})
}
