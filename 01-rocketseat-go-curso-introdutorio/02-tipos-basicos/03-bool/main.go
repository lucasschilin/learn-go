package main

import "fmt"

func main() {
	// Por padrão a variável bool é false
	var bool_value bool
	fmt.Println("Variável bool por padrão é:", bool_value)

	bool_value = true
	fmt.Println("Novo valor da variável bool é:", bool_value)

	// Posso usar expressões para definir o valor de tru or false
	bool_value = 10 < 5
	fmt.Println("10 < 5?", bool_value)

	bool_value = 10 > 5
	fmt.Println("10 > 5?", bool_value)
}
