package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	println("Hello, world!")
	fmt.Println("Hello, world!")
	println(uuid.New().String())

	// Declaracion de constante con tipo de dato
	const pi float64 = 3.14
	fmt.Println(pi)

	// Declaracion de constante sin tipo de dato (Go por defecto ya detecta el tipo de dato)
	const pi2 = 3.1415
	fmt.Println(pi2)

	// Concatenacion
	fmt.Println("Pi:", pi)
	fmt.Println("Pi2:", pi2)
	fmt.Println("Pi:", pi, "- Pi2:", pi2)

	// Variable
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Tipo de datos

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Variables Arimeticas

	x := 10
	y := 50

	result := x + y
	fmt.Println("Suma:", result)

	result = x - y
	fmt.Println("Resta:", result)

	result = x * y
	fmt.Println("Multiplicar:", result)

	result = y / x
	fmt.Println("Dividir:", result)

	// Printf

	nombre := "Platzi"
	cursos := 500

	/*
		%s es para definir un valor en String
		%d es para definir un valor en entero
		%v es para definir cualquier tipo de valor
	*/
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	// Sprintf

	message := fmt.Sprintf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Println(message)
}
