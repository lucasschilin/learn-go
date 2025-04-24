package main

import (
	"errors"
	"fmt"
)

func main() {
	nota := 98

	fmt.Print("Fez a prova e tirou ", nota, ", ... ")

	if nota >= 90 {
		fmt.Println("Muito bom")
	} else if nota >= 70 {
		fmt.Println("Ok")
	} else {
		fmt.Println("Rodou")
	}

	// declaração curta em if. neste caso err só fica disponível dentro do escopo do if
	if err := thisIsAnError(); err != nil {
		fmt.Println(err.Error())
	}

	players := map[string]float64{
		"lucas":  8,
		"amanda": 9.87,
		"bruno":  0,
	}

	if value, ok := players["lucas"]; ok {
		fmt.Println("Pontos do Lucas:", value, ok)
	}

	value, ok := players["bruno"]
	fmt.Println(value, ok)

}

func thisIsAnError() error {
	return errors.New("Este é um ERRO!")
}
