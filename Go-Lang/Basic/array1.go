package main

import "fmt"

func main() {
	var array1 = [3]int{1, 2, 3}
	var array2 = [...]int{10, 20, 30}
	var array3 = [3]string{"nasi goreng", "mie goreng", "soto"}
	array4 := [3]int{1, 2, 3}
	fmt.Println("array1", array1)
	fmt.Println("array2", array2)
	fmt.Println("array3", array3)
	fmt.Println("array4", array4)

	//access element
	var a int = array1[0]
	fmt.Println("a", a)
	//Change Elements of an Array
	array4[2] = 30
	fmt.Println("array4", array4)
}
