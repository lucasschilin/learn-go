package main

import (
	"fmt"
	"time"
)

func showMessage() {
	fmt.Println("Showing message")
}

func main() {
	go showMessage()
	go showMessage()
	go showMessage()
	go showMessage()
	go showMessage()
	go showMessage()
	go showMessage()
	go showMessage()
	go showMessage()
	go showMessage()
	go showMessage()

	fmt.Println("Finalizando")
	time.Sleep(1 * time.Second)
}
