package main

import (
	"context" // Needed if you want to return structured errors
	"fmt"

	"github.com/go-openapi/loads"
	"github.com/wailsapp/wails/v2/pkg/menu"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// func (a *App) handleAbout(data *menu.CallbackData) {
// 	rt.MessageDialog(a.ctx, rt.MessageDialogOptions{
// 		Title:   "关于",
// 		Message: `SepcView 是一款基于 ReDoc 开源的现代化 OpenAPI 文档查看工具，专为开发者、API 设计师和技术文档团队打造。它提供了简洁直观的界面，让您能够轻松浏览和理解 API 规范。`,
// 	})
// }

func (a *App) handleFileOpen(data *menu.CallbackData) {

	selection, err := rt.OpenFileDialog(a.ctx, rt.OpenDialogOptions{
		Title: "Select File",
		Filters: []rt.FileFilter{
			{
				DisplayName: "OpenAPI Files (*.yaml, *.yml, *.json)",
				Pattern:     "*.yaml;*.yml;*.json",
			},
			{
				DisplayName: "All Files (*.*)",
				Pattern:     "*.*",
			},
		},
	})

	if err != nil {
		fmt.Printf("error opening file dialog: %v\n", err)
	}

	rt.EventsEmit(a.ctx, "fileSelected", selection)

}

// SelectOpenFile prompts the user to select an OpenAPI file (YAML or JSON).
// Returns the selected file path or an empty string if cancelled.
func (a *App) SelectOpenFile() (string, error) {
	selection, err := rt.OpenFileDialog(a.ctx, rt.OpenDialogOptions{
		Title: "Select OpenAPI File",
		Filters: []rt.FileFilter{
			{
				DisplayName: "OpenAPI Files (*.yaml, *.yml, *.json)",
				Pattern:     "*.yaml;*.yml;*.json",
			},
			{
				DisplayName: "YAML Files (*.yaml, *.yml)",
				Pattern:     "*.yaml;*.yml",
			},
			{
				DisplayName: "JSON Files (*.json)",
				Pattern:     "*.json",
			},
			{
				DisplayName: "All Files (*.*)",
				Pattern:     "*.*",
			},
		},
	})
	if err != nil {
		// Don't return error for cancellation, just empty path
		return "", fmt.Errorf("error opening file dialog: %w", err)
	}
	// If user cancelled, selection will be empty string ""
	return selection, nil
}

// LoadAndParseOpenAPI loads the spec file from the given path,
// parses it using go-openapi, and returns the spec as a JSON string.
// Returns an error string if parsing fails.
func (a *App) LoadAndParseOpenAPI(filePath string) (string, error) {
	if filePath == "" {
		return "", fmt.Errorf("no file path provided")
	}

	// Use go-openapi/loads to parse the spec.
	// It handles both YAML and JSON automatically.
	specDoc, err := loads.Spec(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to load or parse spec '%s': %w", filePath, err)
	}

	// Return the raw spec as a JSON string
	return string(specDoc.Raw()), nil
}
