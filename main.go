package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	appMenu := menu.NewMenu() // Create an empty menu
	appMenu.Append(menu.AppMenu())

	fileMenu := appMenu.AddSubmenu("配置")
	fileMenu.Append(menu.Text("打开文件", nil, app.handleFileOpen))
	// helpMenu := appMenu.AddSubmenu("关于")
	// helpMenu.Append(menu.Text("关于 My App", nil, app.handleAbout))

	err := wails.Run(&options.App{
		Title:  "SepcView",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []any{
			app,
		},
		Menu:  appMenu,
		Debug: options.Debug{OpenInspectorOnStartup: true},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
