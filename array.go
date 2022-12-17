package main

import "fmt"

func main(){

	var bahasa [5]string
	bahasa[0] = "PHP"
	bahasa[1] = "Javascript"
	bahasa[2] = "Golang"
	bahasa[3] = "C++"
	bahasa[4] = "Java"

	fmt.Println(bahasa)
	fmt.Println("Panjang Array: ", len(bahasa))

	names := [3]string{"Ilham Budiawan", "Bela Nabila", "Dila Azzura"} 
	fmt.Println(names)

	classes := [...]string{
		"TI-P2101",
		"TI-P2102",
		"TI-P2103",
		"TI-P2104",
		"TI-P2105",
	}
	fmt.Println(classes)
	fmt.Println("Panjang Array: ", len(classes))

	for index, class := range classes{
		fmt.Println("Index ke ", index, "Kelas ", class)
	}
}