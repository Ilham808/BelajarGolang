package main

import "fmt"

func main(){
	var myMap map[string]int
	myMap = map[string]int{}

	myMap["satu"] = 1
	myMap["dua"] = 2
	myMap["tiga"] = 3
	myMap["empat"] = 4

	fmt.Println(myMap)

	names := map[string]string{
		"satu":"Ilham Budiawan",
		"dua":"Bela Nabila",
		"tiga":"Dila Azzura",
	}

	fmt.Println(names)

	for index, name := range names{
		fmt.Println("Index ke", index, "Value:", name)
	}

	fmt.Println("=======================")

	delete(myMap, "satu")

	for key, value := range myMap{
		fmt.Println("Index ke", key, "Value:", value)
	}

	fmt.Println("=======================")
	fmt.Println("==== Cek Value Map ====")

	cekValue, isError := names["empat"]
	fmt.Println(cekValue)
	fmt.Println(isError)
}