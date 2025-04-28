package main

import "fmt"

func main() {
	// Um array é imnutável e possui tamanho fixo.
	var gavetas [3]string
	gavetas[0] = "copos"
	gavetas[1] = "talheres"
	gavetas[2] = "panos"

	fmt.Println(gavetas)
	fmt.Println("Copos:", gavetas[0])

}
