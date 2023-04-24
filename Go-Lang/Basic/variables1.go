// DECLEAR VARIABLE

package main

import "fmt"

var x string
var y string = "ini adalah y diluar fungsi"

func main() {
	// Variable Declaration With Initial Value
	var student1 string = "ramma"
	var student2 = "hayu"
	student3 := "fitra"

	fmt.Println("variable1 :", student1)
	fmt.Println("variable2 :", student2)
	fmt.Println("variable3 :", student3)

	// Variable Declaration Without Initial Value
	var a string
	var b int
	var c bool

	fmt.Println("variable a:", a)
	fmt.Println("variable b:", b)
	fmt.Println("variable c:", c)

	// Value Assignment After Declaration
	var student4 string
	student4 = "ramma hayu fitra"
	fmt.Println("student1 :", student4)

	// Difference Between var and :=
	// var
	x = "ini adalah x didalam fungsi"
	fmt.Println("X", x)
	fmt.Println("Y", y)
	// fmt.Println("z", z)
	// untuk z tidak dapat di declear diluar fungsi

}
