package main

import "fmt"

func main() {
	soma := soma(5, 5)
	fmt.Println(soma)

	var fixo = 2
	multiplica := func(a int) int {
		return a * fixo
	}

	resultado := multiplica(15)
	fmt.Println(resultado)
}

func soma(x, y int) int {
	return x + y
}
