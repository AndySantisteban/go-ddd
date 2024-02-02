package main

import (
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.con/AndyGo/go-ddd/cmd/go-app/controllers"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	allnote := controllers.NewControllerAllNote()

	// Create application with options
	err := wails.Run(&options.App{

		Title:  "go-app",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			allnote.Startup(ctx)
		},
		Bind: []interface{}{
			allnote,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
