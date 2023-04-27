package main
import ("fmt")

func main(){
	a := map[string]int{"one":1, "two":2, "three":3, "four":4}

	var b = [] string{} 
	b = append(b, "one", "two", "three", "four")

	for k, v := range a {
		fmt.Println(k, v)
	}

	fmt.Println() 

	for _, element := range b {
		fmt.Println(element, a[element])
	}
}