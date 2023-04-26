// nested if
package main

import "fmt"

func main() {
	x := 20
	if x >= 10 {
		fmt.Println("syarat 1 terpenuhi")
		if x > 15 {
			fmt.Println("syarat 2 terpenuhi")
		}
	} else {
		fmt.Println("kedua sayart tidak terpenuhi")
	}
}
