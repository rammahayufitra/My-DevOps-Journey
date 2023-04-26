package main

import "fmt"

func main() {
	var a = map[string]string{"nama": "ramma", "age": "33"}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["age"] = "30"
	fmt.Println(a)
	fmt.Println(b)

}
