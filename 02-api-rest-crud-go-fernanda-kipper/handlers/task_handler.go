package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (th *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var taskID int

	query := "SELECT id FROM tasks WHERE id = $1;"
	err = th.DB.QueryRow(query, id).Scan(&taskID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Task not found.", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query = "UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4 RETURNING *"
	err = th.DB.QueryRow(query, task.Title, task.Description, task.Status, taskID).
		Scan(&task.ID, &task.Title, &task.Description, &task.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (th *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var taskID int

	query := "SELECT id FROM tasks WHERE id = $1;"
	err = th.DB.QueryRow(query, id).Scan(&taskID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Task not found.", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query = "DELETE FRom tasks WHERE id = $1;"
	_, err = th.DB.Exec(query, taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task deleted successfully."))

}
