package main

import "fmt"

func main() {

	var float_32_bits float32 = 1.8
	var float_64_bits float64 = 3.14

	//Não é possível fazer multiplicação dois tipos diferentes
	// multi := float_32_bits * float_64_bits
	multi := float_32_bits * float32(float_64_bits)

	fmt.Println("float_32_bits: ", float_32_bits)
	fmt.Println("float_64_bits: ", float_64_bits)

	fmt.Println(multi)
}
