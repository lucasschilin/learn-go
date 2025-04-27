package main

import "fmt"

type Pessoa struct {
	Nome string
}

func main() {
	var p1 Pessoa = Pessoa{Nome: "Lucas"}
	var p2 Pessoa = Pessoa{Nome: "Amanda"}

	var p3 *Pessoa = &p1

	//p1.Nome e p3.Nome estão no mesmo endereço
	fmt.Println(&p1.Nome, &p3.Nome)

	fmt.Println(p1, p3)

	p3.Nome = "Matheus"
	fmt.Println(p1, p3)

	var p4 Pessoa = p2
	fmt.Println(p2, p4)

	p4.Nome = "Márcia"
	fmt.Println(p2, p4)
}
