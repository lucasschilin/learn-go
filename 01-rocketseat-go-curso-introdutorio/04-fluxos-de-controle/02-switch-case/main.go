package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Quando é sexta?")
	today := time.Now().Weekday()

	switch time.Friday {
	case today:
		fmt.Println("Hoje")
	case today + 1:
		fmt.Println("Amanhã")
	case today + 2:
		fmt.Println("Depois de amanhã")
	default:
		fmt.Println("Sem previsão kk")
	}
}
