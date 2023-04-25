package main

import "fmt"

func main() {
	slice1 := []string{}
	slice2 := []string{"Yamaha", "Honda", "Suzuki", "Ducati", "BMW"}
	fmt.Println(len(slice1), len(slice2))
	fmt.Println(cap(slice1), cap(slice2))
	fmt.Println(slice1, slice2)

	// create slice from array
	var array1 = [...]int{1, 2, 3, 4, 5, 6, 7}
	slice3 := array1[0:4]
	fmt.Println(len(slice3), cap(slice3), slice3)
	// Create a Slice With The make() Function

	var slice4 = make([]int, 5, 10)
	var slice5 = make([]int, 5)
	fmt.Println(len(slice4), cap(slice4), slice4)
	fmt.Println(len(slice5), cap(slice5), slice5)
}
