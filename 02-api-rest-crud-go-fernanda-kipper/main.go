package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lucasschilin/learn-go/apirest-crud-go/config"
	"github.com/lucasschilin/learn-go/apirest-crud-go/handlers"
	"github.com/lucasschilin/learn-go/apirest-crud-go/models"
)

func main() {
	dbConnection := config.SetupDB()
	defer dbConnection.Close()

	_, err := dbConnection.Query(models.CreateTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	taskHandler := handlers.NewTaskHandler(dbConnection)

	router.HandleFunc("/tasks", taskHandler.PostTask).Methods("POST")
	router.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	// router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("UPDATE")
	// router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
