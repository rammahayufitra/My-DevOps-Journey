package main

import "fmt"

func main() {
	const A int = 1
	fmt.Println("A", A)
	const B = 2
	fmt.Println("B", B)
	const (
		C int    = 1
		D        = 3.14
		F string = "Hello"
	)
	fmt.Println("C", C, "D", D, "F", F)

}
