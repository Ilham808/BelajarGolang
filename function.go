package main

import "fmt"

func main(){
	helloWorld()
	nameStudent("Ilham Budiawan", 20)
	fmt.Println("5 + 8:", count(5,8))
}

func helloWorld(){
	fmt.Println("Hello World")
}

func nameStudent(name string, age int){
	fmt.Println("Hallo, nama saya ", name, "Umur saya", age, "tahun")
}

func count(numberOne int, numberTwo int) int{
	return numberOne + numberTwo
}