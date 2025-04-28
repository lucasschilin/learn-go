package meet

import "fmt"

/*
Função/variável iniciada com letra minúscula só fica visível
internamente para o pacote, neste caso, o package "meet"
*/
func say(word string) {
	fmt.Println(word)
}
