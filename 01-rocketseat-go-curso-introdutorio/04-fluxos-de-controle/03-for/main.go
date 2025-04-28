package main

import (
	"fmt"
)

func main() {
	var sum int
	fmt.Println("FOR")
	for i := 1; i <= 20; i++ {
		sumi := sum + i
		fmt.Println(sum, "+", i, "=", sumi)

		sum = sumi
	}

	//while (posso simular o while)
	fmt.Println("\n'WHILE'")
	cont := 0
	for cont <= 20 {
		fmt.Println("cont:", cont)

		cont++
	}

	//percorrendo slice
	fmt.Println("\nPERCORENDO SLICE")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(slice); i++ {
		fmt.Println("índice", i, ": ", slice[i])
	}

}
