package main

import (
	"fmt"
	"strings"
)

func main() {
	var_ola := "Olá"

	var var_mundo string = "Mundo"

	var_ola_mundo := var_ola + " " + var_mundo

	fmt.Println(var_ola_mundo)
	fmt.Println("TO UPPER:", strings.ToUpper(var_ola_mundo))
	fmt.Println("to lower:", strings.ToLower(var_ola_mundo))
	fmt.Println("Contains 'Mun':", strings.Contains(var_ola_mundo, "Mun"))
	fmt.Println("Compare:", strings.Compare(var_ola, "Olááá"))

}
