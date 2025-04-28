package main

import "fmt"

func main() {
	// Inteiro de 8 bits (Valores de -128 a 127)
	var int_8_bits int8 = 127

	// Inteiro de 16 bits (Valores de -32768 a 32767)
	var int_16_bits int16 = 12000

	// Inteiro de 32 bits (Valores de -2147483648 a 2147483647)
	var int_32_bits int32 = -54651324

	// Inteiro de 64 bits
	// (Valores de -9223372036854775808 a 9223372036854775807)
	var int_64_bits int64 = 87456464321

	// Caso não for definido, pega por padrão a arquitetura (neste caso o 64)
	var int_default int = 132433

	fmt.Println(int_8_bits, int_16_bits, int_32_bits, int_64_bits, int_default)
}
