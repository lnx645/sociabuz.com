package utils

import (
	"embed"
	"fmt"
	"html/template"
	"sync"

	"github.com/goccy/go-json"
)

type ViteAssets struct {
	JS  string
	CSS []string
}
type ManifestEntry struct {
	File    string   `json:"file"`
	Name    string   `json:"name,omitempty"`
	Src     string   `json:"src,omitempty"`
	IsEntry bool     `json:"isEntry,omitempty"`
	CSS     []string `json:"css,omitempty"`
	Assets  []string `json:"assets,omitempty"`
}

var (
	manifest map[string]ManifestEntry
	mu       sync.RWMutex
)

func LoadManifest(manifestFile embed.FS) {
	data, err := manifestFile.ReadFile("public/manifest.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(data, &manifest); err != nil {
		panic("WKWKW")
	}
}

func AssetsViteResolver(asset string) ViteAssets {

	mu.RLock()
	entry, ok := manifest[asset]
	defer mu.RUnlock()
	if !ok {
		return ViteAssets{}
	}

	return ViteAssets{
		JS:  entry.File,
		CSS: entry.CSS,
	}

}

func RenderViteSource(source string) template.HTML {
	var html string
	viteUrl := GetEnv("VITE_URL")
	assetsURL := GetEnv("ASSETS_URL")
	if GetEnv("SERVER_ENV") != "production" {
		html += fmt.Sprintf(`<script src="%s/@vite/client" type="module"></script>`, viteUrl)
		html += "\n"
		html += fmt.Sprintf(`<script src="%s/src/main.ts" type="module"></script>`, viteUrl)
	} else {
		data := AssetsViteResolver(source)
		if data.JS != "" {
			html += fmt.Sprintf(`<script src="%s/%s" type="module"></script>`, assetsURL, data.JS)
			for _, css := range data.CSS {
				html += "\n"
				html += fmt.Sprintf(`<link rel="stylesheet" href="%s/%s"></link>`, assetsURL, css)
			}
		}
	}
	return template.HTML(html)
}
