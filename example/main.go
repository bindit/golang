package main

import "emath"
import "fmt"

func main() {
	testFibonnaci()
	testString()
}

func testFibonnaci() {
	f := emath.CalculateOne(10);
	fmt.Printf("%v", f)
}

func testString() {
	fmt.Printf("to do")
}