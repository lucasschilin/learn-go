package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/lucasschilin/learn-go/apirest-crud-go/models"
)

type TaskHandler struct {
	DB *sql.DB
}

func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

func (th *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM tasks;"
	rows, err := th.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []models.Task = []models.Task{}
	for rows.Next() {
		var task models.Task

		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (th *TaskHandler) PostTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO tasks 
					(title, description, status) 
				VALUES 
					($1, $2, $3) 
				RETURNING id, title, description, status;`
	err = th.DB.
		QueryRow(query, newTask.Title, newTask.Description, newTask.Status).
		Scan(&newTask.ID, &newTask.Title, &newTask.Description, &newTask.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}
