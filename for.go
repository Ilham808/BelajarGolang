package main

import "fmt"

func main(){
	for i:=0; i <= 5; i++{
		fmt.Println("For ke-1 Angka ke ", i)
	}

	one := 1
	for one <= 5{
		fmt.Println("For ke-2 Angka ke: ", one)
		one++
	}

	title := "Belajar Golang"
	for index, letter := range title{
		fmt.Println("Index :", index, " letter :", string(letter))
	}
}