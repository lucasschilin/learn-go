package models

import (
	"fmt"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

const TableName string = "tasks"

var CreateTableQuery = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	id SERIAL PRIMARY KEY,
	title VARCHAR(100) NOT NULL,
	description TEXT NULL,
	status BOOLEAN NOT NULL DEFAULT FALSE
);`, TableName)
