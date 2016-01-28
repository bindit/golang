package main

import "emath"
import "fmt"
import "string"

func main() {
	testFibonnaci()
	testString()
}

func testFibonnaci() {
	f := emath.CalculateOne(10);
	fmt.Printf("%v", f)
}

func testString() {
	fmt.Printf(string.Reverse("Hello World, dawg"))
}