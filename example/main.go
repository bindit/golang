package main

import "fibonacci"
import "fmt"

func main() {
	f := fibonacci.CalculateOne(10);
	fmt.Printf("%v", f)
}
