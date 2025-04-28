package main

import "fmt"

func main() {
	var portas []string
	fmt.Println(portas)

	portas = append(portas, "quarto")
	portas = append(portas, "sala")
	portas = append(portas, "cozinha")

	fmt.Println("Todas as portas:", portas)

	qtde_portas := len(portas)

	fmt.Println("Qtde. de portas:", len(portas))
	fmt.Println("Primeira porta:", portas[0])
	fmt.Println("Última porta:", portas[qtde_portas-1])

	//divisão de slice (slice[x:y+1]),
	// onde x é o index do elemento inicial e y é o index do elemento final
	fmt.Println("Últimas duas portas: ", portas[1:3], "ou", portas[1:])
	fmt.Println("Primeiras duas portas: ", portas[0:2], "ou", portas[:2])

	portas = portas[1:3]
	fmt.Println(portas)

}
