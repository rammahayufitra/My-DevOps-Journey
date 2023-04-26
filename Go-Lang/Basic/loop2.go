package main

// continue & break
import "fmt"

func main() {
	for x := 0; x < 100; x++ {
		if x == 2 {
			// continue
			break
		}
		fmt.Println(x)
	}
}
