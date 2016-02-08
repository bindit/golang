package main

import (
	"object/car"
	"emath"
	"estring"
	"fmt"
)

func main() {
	testFibonnaci()
	testString()
	testObject()
	testMultipleReturn()
	testMultipleArgs()
	testGenerator()
}

func testFibonnaci() {
	f := emath.Calculate(10);

	for _,number := range f {
		fmt.Printf("%d\n", number)
	}
}

func testString() {
	fmt.Printf("%s\n", estring.Reverse("Hello World, dawg"))
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

func testMultipleArgs() {
	passMultipleInt(1,2,3)

	slice := []int{1,2,3}
	passMultipleInt(slice...)

	arr := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	passMutlipleInterface(1, "test", 0.2, arr)
}

func passMultipleInt(args ...int)  {
	var total int
	for _,v := range args {
		total += v
	}

	fmt.Printf("total: %d\n", total)
}

func passMutlipleInterface(args ...interface{})  {
	for _,v := range args {
		fmt.Printf("%v\n", v)
	}
}

func testMultipleReturn() {
	int, int2 := returnTwo()
	fmt.Printf("%d, %d\n", int, int2)

	int, int2, float := returnThree()
	fmt.Printf("%d, %d, %f\n", int, int2, float)
}

func returnTwo() (int, int) {
	return 1, 2
}

func returnThree() (int, int, float64)  {
	return 1, 2, 0.2
}

func printCarSpeed(car car.Rider)  {
	fmt.Printf("%v\n", car.Speed())
}

func testGenerator()  {
	generator := makeGenerator()
	fmt.Printf("%d\n", generator())
	fmt.Printf("%d\n", generator())
	fmt.Printf("%d\n", generator())
}

func makeGenerator() func() int {
	i := 0

	return func() int {
		i += 2
		return i
	}
}