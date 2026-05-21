package config

import (
	"embed"
	"io/fs"
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/django/v4"
	"sobuz.id/application/app/utils"
)

func FiberConfig(views embed.FS) fiber.Config {

	sub, err := fs.Sub(views, "views")

	if err != nil {
		panic("Failed to embed views folder: " + err.Error())
	}

	engine := django.NewFileSystem(http.FS(sub), ".html")
	engine.AddFunc("vite", func(path string) utils.ViteAssets {
		if utils.GetEnv("SERVER_ENV") != "production" {
			return utils.ViteAssets{
				JS:  utils.GetEnv("VITE_URL") + "/" + path,
				CSS: nil,
			}
		}
		return utils.AssetsViteResolver(path)
	})
	return fiber.Config{
		Views:         engine,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		StrictRouting: true,
	}
}
