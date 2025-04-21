package meet

import "fmt"

/*
Quando função inicia com letra maiuscula, ela SÓ fica visível
internamente para o módulo, neste caso, o módulo "meet"
*/

func say(word string) {
	fmt.Println(word)
}
