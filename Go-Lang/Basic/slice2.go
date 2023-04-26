package main
import "fmt" 

func main(){
	slice1 := []int{10,20,30}
	fmt.Println("slice1", slice1)
	slice1[2] = 50
	fmt.Println("slice1", slice1)
	slice1 = append(slice1, 40,50)
	fmt.Println("slice1", slice1)
	
	numbers := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("numbers", numbers)
	fmt.Println("length", len(numbers))
	fmt.Println("capacity", cap(numbers))

	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers) 
	fmt.Println("neededNumbers", neededNumbers)
	fmt.Println("NumberCopy", numbersCopy)
	fmt.Println("lenght", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

}