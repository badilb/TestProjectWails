package backend

import (
	"WailsToDoList/backend/dao"
	"WailsToDoList/backend/models"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Инициализация DAO
var dataAccess *dao.DAO

// InitializeDatabase устанавливает соединение с базой данных и создает DAO
func InitializeDatabase() error {
	var err error
	connStr := "user=postgres dbname=todoapp sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	dataAccess = dao.NewDAO(db)
	return nil
}

// CreateTaskHandler обрабатывает создание новой задачи
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := dataAccess.CreateTask(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// GetTasksHandler возвращает список задач
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := dataAccess.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// UpdateTaskHandler обрабатывает обновление задачи
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := mux.Vars(r)["id"]

	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := dataAccess.UpdateTask(taskID, updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTask)
}

// DeleteTaskHandler обрабатывает удаление задачи
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := mux.Vars(r)["id"]

	if err := dataAccess.DeleteTask(taskID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
