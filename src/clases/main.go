package main

import (
	"fmt"

	pk "github.com/Relari/myFirstGo/src/clases/mypackage"
)

// Clases privados
type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	var otherCard car
	otherCard.brand = "Ferrari"
	fmt.Println(otherCard)

	var myCarOther pk.CarPublic
	myCarOther.Brand = "Toyota"

	fmt.Println(myCarOther)

	fmt.Println(pk.PrintMessage("Hello World"))
}
