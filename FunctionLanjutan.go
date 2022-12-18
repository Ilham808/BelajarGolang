package main

import "fmt"

func main(){
	numbers := []int{10,49,2,5,6}
	result := sum(numbers)
	fmt.Println(result)
}

func sum(numbers []int) int{
	var total int
	for _, number := range numbers{
		total = total + number
	}
	return total
}