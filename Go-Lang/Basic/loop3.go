// nested loop

package main

import "fmt"

func main() {
	size := [2]string{"big", "small"}
	var fruits = [3]string{"apple", "banana", "cerry"}
	for i := 0; i < len(size); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(size[i], fruits[j])
		}
	}
}
