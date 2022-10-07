package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer Map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar valor

	valor, ok := m["Jose"]
	fmt.Println(valor, ok)

	valor2, ok := m["Josep"]
	fmt.Println(valor2, ok)
}
