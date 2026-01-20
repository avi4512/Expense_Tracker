package main

import "fmt"

func main() {

	fmt.Println("Hello World...")
	add(10, 20)
}

func add(x, y int) {
	res := x + y
	fmt.Println("Result = ", res)
}
