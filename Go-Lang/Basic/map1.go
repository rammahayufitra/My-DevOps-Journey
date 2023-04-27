package main

import "fmt"

func main() {
	var a = map[string]string{"nama-depan": "ramma", "nama-tengah": "hayu", "nama-belakang": "fitra"}
	b := map[string]int{"juara1": 1, "juara2": 2, "juara3": 3}
	var c = make(map[string]string)
	c["makanan"] = "nasi goreng"
	c["minuman"] = "Teh Openg"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(c["makanan"])
	c["makanan"] = "soto"
	fmt.Println(c)
	delete(c, "makanan")
	fmt.Println(c)
}
