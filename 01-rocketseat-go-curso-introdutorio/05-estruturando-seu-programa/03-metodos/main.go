package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Apresentar() {
	fmt.Println("Olá a todos. Este é", p.Nome, "e possui", p.Idade, "anos! ")
}

func main() {

	p1 := Pessoa{
		Nome:  "Schilin",
		Idade: 24,
	}

	fmt.Println(p1)
	p1.Apresentar()
}
