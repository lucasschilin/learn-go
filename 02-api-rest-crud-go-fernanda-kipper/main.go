package main

import (
	"fmt"

	"github.com/lucasschilin/learn-go/apirest-crud-go/config"
)

func main() {
	dbConnection := config.SetupDB()

	fmt.Println(dbConnection)
}
