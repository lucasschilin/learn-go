package main

import (
	"fmt"
	"strings"
)

func main() {

	//range em slice
	nums := []int{10, 20, 30, 40, 50}

	for key, val := range nums {
		fmt.Println(key, "->", val)
	}

	//range em map
	people := map[string]int{
		"lucas":   24,
		"amanda":  24,
		"luiz":    57,
		"matheus": 15,
	}

	for person_name, age := range people {
		fmt.Println(strings.ToTitle(person_name), "está com", age, "anos.")
	}

	// posso pegar só akey ou só o val utilizando o "_"
	for _, val := range people { //key, _
		fmt.Println(val)
	}

}
