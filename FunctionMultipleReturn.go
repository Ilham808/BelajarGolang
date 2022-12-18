package main

import "fmt"

func main(){
	luas, keliling := calculate(30, 42)
	fmt.Println("Hasil Luas:" , luas)
	fmt.Println("Hasil Keliling:", keliling)
}

func calculate(panjang int, lebar int) (int, int){
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}