package main

import "fmt"

func callFunction() {
	fmt.Println("call function")
}

func functionWithParameter(fname string, age int) {
	fmt.Println("nama :", fname, "usia :", age)
}

func functionSingleReturnValue(x int, y int) (result int) {
	result = x + y
	return
}

func functionMultipleReturnValue(x int, y string) (angka int, word string) {
	angka = x * x
	word = y
	return
}

func functionRecursion(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return functionRecursion(x + 1)

}

func main() {
	callFunction()
	functionWithParameter("ramma", 33)
	result := functionSingleReturnValue(10, 15)
	fmt.Println(result)
	angka, word := functionMultipleReturnValue(56, "hallo")
	fmt.Println(angka, word)
	functionRecursion(0)
}
