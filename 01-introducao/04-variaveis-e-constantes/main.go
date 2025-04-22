package main

import "fmt"

const city string = "Feliz"
const state = "RS"

func main() {
	//Variável precisa ser tipada
	var name string = "Lucas"

	//Se não for tipada recebe o tipo do valor inicial
	var lastname = "Schlindwein"

	//Com inicialização curta recebe o tipo do valor inicial
	age := 24

	//Não é possível modificar constante
	// city = "POA"

	fmt.Println(name + " " + lastname)
	fmt.Println(age)
	fmt.Println(city + "/" + state)
}
