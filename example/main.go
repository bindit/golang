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
		fmt.Printf("%d", number)
	}
}

func testString() {
	fmt.Printf(string.Reverse("Hello World, dawg"))
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
	fmt.Printf("%v", car.Speed())
}