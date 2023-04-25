// Array Initialization

package main

import "fmt"

func main() {
	var array1 = [5]int{}                            // not
	var array2 = [5]int{1, 2}                        //partial
	var array3 = [5]int{1, 2, 3, 4, 5}               //fully
	var array4 = [5]int{1: 10, 2: 40}                // Initialize Only Specific Elements
	var array5 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9} //Find the Length of an Array

	fmt.Println("array1", array1)
	fmt.Println("array2", array2)
	fmt.Println("array3", array3)
	fmt.Println("array4", array4)
	fmt.Println("array5", len(array5))
}
