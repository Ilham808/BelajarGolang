package main

import "fmt"

func main(){
	students := []map[string]string{
		{"name":"Ilham Budiawan", "class":"TI-P2101"},
		{"name":"Bela Nabila", "class":"TI-P2102"},
		{"name":"Dila Azzura", "class":"TI-P2103"},
	}
	fmt.Println(students)
	fmt.Println("=======================")

	for _, student := range students{
		for index, value := range student{
			fmt.Println("Index: ", index)
			fmt.Println("Value: ", value)
		}
	}

	fmt.Println("=======================")

	for _, studentTwo := range students{
		fmt.Println("Nama ", studentTwo["name"], "Kelas ", studentTwo["class"])
	}
}