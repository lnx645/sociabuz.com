package config

import (
	"embed"
	"io/fs"
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v3"
	"sobuz.id/application/app/utils"
)

func FiberConfig(views embed.FS) fiber.Config {

	sub, err := fs.Sub(views, "views")

	if err != nil {
		panic("Failed to embed views folder: " + err.Error())
	}

	engine := html.NewFileSystem(http.FS(sub), ".html")
	engine.AddFunc("vite_render", utils.RenderViteSource)

	return fiber.Config{
		Views:         engine,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		StrictRouting: true,
	}
}
