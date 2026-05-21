package infra

import (
	"embed"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/extractors"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/gofiber/fiber/v3/middleware/session"
	"sobuz.id/application/app/config"
	"sobuz.id/application/routes"
)

type ServerConfig struct {
	ViewsDir     embed.FS
	PublicDir    embed.FS
	ManifestFile embed.FS
}

type ServerApp struct {
	App    *fiber.App
	config *config.Config
}

func InitHTTPServer(cfg *ServerConfig) ServerApp {
	fiberCfg := config.FiberConfig(cfg.ViewsDir)
	cookieSession := session.NewStore(session.Config{
		CookieSecure: true,
	})
	//fiber configuration
	app := fiber.New(fiberCfg)
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(helmet.New())
	app.Use(csrf.New(csrf.Config{
		CookieName:        "_csrf_token",
		CookieSecure:      true,
		CookieHTTPOnly:    true,
		CookieSessionOnly: true,
		CookieSameSite:    "lax",
		Extractor: extractors.Chain(
			extractors.FromForm("_csrf"),
			extractors.FromHeader("X-Csrf-Token"),
		),
		Session: cookieSession,
	}))
	//file server configuration
	routes.AssetsDirectoryInitialize(app, cfg.PublicDir)
	return ServerApp{App: app, config: config.AppConfig}
}

func (s *ServerApp) Shutdown() error {
	return s.App.Shutdown()
}

func (s *ServerApp) Run() error {
	url := fmt.Sprintf("%s:%s", s.config.ServerHost, s.config.ServerPort)
	fmt.Println(url)
	return s.App.Listen(url)
}
