package utils

import (
	"os"
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
	once     sync.Once
	mu       sync.RWMutex
)

func LoadManifest() {
	data, err := os.ReadFile("public/manifest.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(data, &manifest); err != nil {
		panic("WKWKW")
	}
}

func AssetsViteResolver(asset string) ViteAssets {

	once.Do(LoadManifest)
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
