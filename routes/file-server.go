package routes

import (
	"embed"
	"io/fs"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func AssetsDirectoryInitialize(app *fiber.App, dir embed.FS) {

	sub, err := fs.Sub(dir, "public")

	if err != nil {
		panic(err)
	}

	app.Get("/*", static.New("", static.Config{
		FS: sub,
	}))
}
