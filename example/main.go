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
	f := emath.CalculateOne(10);
	fmt.Printf("%v", f)
}

func testString() {
	fmt.Printf(string.Reverse("Hello World, dawg"))
}

func testObject() {
	c := object.Car {
		Color: "Red",
		Engine: object.Engine {
			Capacity: 1.8,
			Power: 120,
		},
	}
	fmt.Printf("%v", c)
}