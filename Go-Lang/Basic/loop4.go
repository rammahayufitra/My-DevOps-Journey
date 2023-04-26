package main

import "fmt"

func main() {
	var menu = [3]string{"nasi goreng", "mie goreng", "bakso"}
	for i, item := range menu {
		fmt.Println(i, item)
	}
}
