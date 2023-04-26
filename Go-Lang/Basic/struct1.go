package main

import "fmt"

type biodata struct {
	name string
	age  int
}

func printResult(data biodata) {
	fmt.Println("name :", data.name)
	fmt.Println("age :", data.age)
}

func main() {
	var data1 biodata
	data1.name = "ramma"
	data1.age = 33

	var data2 biodata
	data2.name = "ramma"
	data2.age = 33

	printResult(data1)
	printResult(data2)

	fmt.Println(data1.name, data1.age)
}
