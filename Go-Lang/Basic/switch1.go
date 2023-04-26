// singlae case
package main

import "fmt"

func main() {
	var data int = 1
	switch data {
	case 1:
		fmt.Println("Found")
	default:
		fmt.Println("Not Found")
	}
}
