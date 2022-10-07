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

	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

	normalFunction("Hello World")
	tripletFunction(1, 2, "message")
	tripletFunction2(1, 2, "message 2")

	fmt.Println("Return Value", returnValue(2))

	valueOne, valueTwo := returnTwoValues(2)
	fmt.Println("Return Value 1 and Value 2", valueOne, valueTwo)

	valueThree, _ := returnTwoValues(3)
	fmt.Println("Return Value 3", valueThree)

	// Ciclo For

	for i := 0; i <= 10; i++ {
		fmt.Println("Value", i)
	}

	for a < 10 {
		fmt.Println("Value a", a)
		a++
	}

	// Bucle Switch
	switch module := 4 % 2; module {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es inpar")
	}

	// Sin condicion
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No es condicion")
	}

	var array [4]int

	array[0] = 1
	array[1] = 2

	fmt.Println(array, len(array), cap(array))

	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice

	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append

	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva lista

	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)

	fmt.Println(slice)
}

// funciones

func normalFunction(message string) {
	fmt.Println(message)
}

func tripletFunction(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func tripletFunction2(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func returnTwoValues(a int) (b, c int) {
	return a, a * 2
}
