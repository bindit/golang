package main

import (
	"object/car"
	"emath"
	"string"
	"fmt"
)

func main() {
	testFibonnaci()
	testString()
	testObject()
}

func testFibonnaci() {
	f := emath.Calculate(10);

	for _,number := range f {
		fmt.Printf("%d\n", number)
	}
}

func testString() {
	fmt.Printf("%s\n", string.Reverse("Hello World, dawg"))
	fmt.Printf("%d\n", "Hello"[1]) //return byte representation
	fmt.Printf("%s\n", "Hello" + " " + "World")
}

func testObject() {
	suv := car.Suv {
		car.Car {
			Color: "Red",
		},
	}
	//suv.Color = "Red"
	suv.Engine = car.Engine {
		Capacity: 1.8,
		Power: 120,
	}

	printCarSpeed(suv)

	coupe := car.Coupe {}
	printCarSpeed(coupe)
}

func printCarSpeed(car car.Rider)  {
	fmt.Printf("%v\n", car.Speed())
}