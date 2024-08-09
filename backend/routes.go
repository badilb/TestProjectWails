package backend

import (
	"github.com/gorilla/mux"
)

// SetupRoutes настраивает маршруты для приложения
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", CreateTaskHandler).Methods("POST")
	router.HandleFunc("/tasks", GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/tasks/{id}", DeleteTaskHandler).Methods("DELETE")

	return router
}
