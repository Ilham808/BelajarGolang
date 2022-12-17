package main

import "fmt"

func main(){
	var names []string

	names = append(names, "Ilham Budiawan")
	names = append(names, "Bela Nabila")
	names = append(names, "Dila Azzura")

	fmt.Println(names)

	classes := []string{
		"TI-P2101",
		"TI-P2102",
		"TI-P2103",
		"TI-P2104",
		"TI-P2105",
	}
	fmt.Println(classes)

	for index, class := range classes{
		fmt.Println("Index ke: ", index, "Kelas: ", class)
	}
}