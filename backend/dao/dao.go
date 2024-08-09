package dao

import (
	"WailsToDoList/backend/models"
	"database/sql"
	_ "github.com/lib/pq"
)

// DAO структура для доступа к данным
type DAO struct {
	DB *sql.DB
}

// NewDAO создает новый DAO с подключением к базе данных
func NewDAO(db *sql.DB) *DAO {
	return &DAO{DB: db}
}

// CreateTask добавляет новую задачу в базу данных
func (dao *DAO) CreateTask(task models.Task) error {
	query := `INSERT INTO tasks (id, title, description, is_completed, created_at) VALUES ($1, $2, $3, $4, NOW())`
	_, err := dao.DB.Exec(query, task.ID, task.Title, task.Description, task.IsCompleted)
	return err
}

// GetTasks получает все задачи из базы данных
func (dao *DAO) GetTasks() ([]models.Task, error) {
	rows, err := dao.DB.Query(`SELECT id, title, description, is_completed, created_at FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// UpdateTask обновляет задачу в базе данных
func (dao *DAO) UpdateTask(taskID string, updatedTask models.Task) error {
	query := `UPDATE tasks SET title=$1, description=$2, is_completed=$3 WHERE id=$4`
	_, err := dao.DB.Exec(query, updatedTask.Title, updatedTask.Description, updatedTask.IsCompleted, taskID)
	return err
}

// DeleteTask удаляет задачу из базы данных
func (dao *DAO) DeleteTask(taskID string) error {
	query := `DELETE FROM tasks WHERE id=$1`
	_, err := dao.DB.Exec(query, taskID)
	return err
}
