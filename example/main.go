package main

import "emath"
import "fmt"
import "string"
import "object"

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
	suv := object.Suv{}
	suv.Color = "Red"
	suv.Engine = object.Engine{
		Capacity: 1.8,
		Power: 120,
	}

	printCarSpeed(suv)

	coupe := object.Coupe{}
	printCarSpeed(coupe)
}

func printCarSpeed(car object.Rider)  {
	fmt.Printf("%v", car.Speed())
}