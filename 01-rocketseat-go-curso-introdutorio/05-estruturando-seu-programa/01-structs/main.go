package main

import "fmt"

type Pessoa struct {
	Nome  string
	Email string
	Idade int
	Mae   Mae
}

type Mae struct {
	Nome  string
	Idade int
}

func main() {
	pessoa1 := Pessoa{
		Nome:  "Lucas Schlindwein",
		Email: "lucasschilin@gmail.com",
		Idade: 24,
		Mae: Mae{
			Nome:  "Marcia",
			Idade: 49,
		},
	}

	var pessoa2 Pessoa
	pessoa2 = Pessoa{
		Nome:  "Amanda Hunoff",
		Email: "amanda@email.com",
		Idade: 24,
	}

	mae2 := Mae{
		Nome:  "Leila",
		Idade: 53,
	}

	pessoa2.Mae = mae2
	fmt.Println(pessoa1, pessoa2)

	pessoa2.Mae.Idade = 52
	fmt.Println(pessoa1, pessoa2)

	pessoa1.Mae = Mae{
		Nome:  "Márcia Janete",
		Idade: 50,
	}

	fmt.Println(pessoa1, pessoa2)
}
