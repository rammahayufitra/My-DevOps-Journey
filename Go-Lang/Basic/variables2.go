// DECLEAR MULTIPLE VARIABLE
package main

import "fmt"

func main() {
	var a, b, c, d int = 1, 2, 3, 4
	fmt.Println("multiple 1 :", a, b, c, d)

	var e, f = 5, "huruf 1"
	g, h := 6, "huruf 2"
	fmt.Println("multiple 2 :", e, f, g, h)

	var (
		i int
		j int    = 1
		k string = "huruf 3"
	)
	fmt.Println("multiple 3 :", i, j, k)

}
