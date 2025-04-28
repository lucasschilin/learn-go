package main

import "fmt"

func main() {
	var pessoas = map[string]int{}

	pessoas["lucas"] = 24
	pessoas["amanda"] = 24
	pessoas["luiz"] = 57

	fmt.Println(pessoas)

	if idade, ok := pessoas["lucas"]; ok {
		fmt.Println("Pessoa existe no map.", idade, ok)
	} else {
		fmt.Println("Pessoa NÃO existe no map.")
	}

	pessoas["lucas"]++

	fmt.Println(pessoas["lucas"])

	delete(pessoas, "lucas")

	if idade, ok := pessoas["lucas"]; ok {
		fmt.Println("Pessoa existe no map.", idade, ok)
	} else {
		fmt.Println("Pessoa NÃO existe no map.")
	}

	delete(pessoas, "diego")

	fmt.Println(pessoas)
}
