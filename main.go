package main

import (
	"context"
	"embed"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"WailsToDoList/backend" // Обновите путь в зависимости от вашего расположения
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Создаем новый роутер
	r := mux.NewRouter()

	// Регистрация маршрутов
	r.HandleFunc("/tasks", backend.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks", backend.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id:[0-9]+}", backend.UpdateTaskHandler).Methods("PUT")
	r.HandleFunc("/tasks/{id:[0-9]+}", backend.DeleteTaskHandler).Methods("DELETE")

	// Запуск HTTP сервера в отдельной горутине
	go func() {
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Инициализация Wails приложения
	err := wails.Run(&options.App{
		Title:  "WailsToDoList",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			// Начальная инициализация
			log.Println("Wails приложение запущено")
		},
		Bind: []interface{}{
			NewApp(), // Привязываем App из app.go
		},
	})

	if err != nil {
		log.Fatalf("Failed to run app: %v", err)
	}
}
