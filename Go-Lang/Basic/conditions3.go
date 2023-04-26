// else if statements
package main

import "fmt"

func main() {
	x := 10
	y := 30

	if x > y {
		fmt.Println("x greater than y")
	} else if x < y {
		fmt.Println("x smaller than y")
	} else {
		fmt.Println("wrong logic")
	}
}
