package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"sobuz.id/application/app/config"
	"sobuz.id/application/app/infra"
	"sobuz.id/application/app/utils"
	"sobuz.id/application/routes"
)

//go:embed views/*
var viewsDIR embed.FS

//go:embed public/*
var publicDIR embed.FS

//go:embed public/manifest.json
var manifestFile embed.FS

func main() {
	utils.LoadManifest(manifestFile)
	quit := make(chan os.Signal, 1)
	//fiber config
	config.InitAppConfig()
	app := infra.InitHTTPServer(&infra.ServerConfig{
		ViewsDir:  viewsDIR,
		PublicDir: publicDIR,
	})

	routes.InitializeRoutes(app.App)
	fmt.Println(config.AppConfig.ServerPort)
	go func() {
		if err := app.Run(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Mematikan server secara perlahan (Graceful Shutdown)...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Error during shutdown: %v", err)
	}
	log.Println("Server Invitnesia berhasil dimatikan dengan aman.")
}
