package main

import "fmt"

func main(){
	age := 18

	if age >= 18{
		fmt.Println("Anda sudah boleh buat KTP")
	}

	averageNilai := 85
	var nilaiHuruf string

	if averageNilai >= 81 && averageNilai <= 100 {
		nilaiHuruf = "A" 
	} else if averageNilai >= 71 && averageNilai <= 80 {
		nilaiHuruf = "B+"
	} else if averageNilai >= 67 && averageNilai <= 70 {
		nilaiHuruf = "B"
	} else if averageNilai >= 61 && averageNilai <= 76 {
		nilaiHuruf = "C+"
	} else if averageNilai >= 56 && averageNilai <= 60 {
		nilaiHuruf = "C"
	} else if averageNilai >= 41 && averageNilai <= 55 {
		nilaiHuruf = "D"
	} else if averageNilai >= 0 && averageNilai <= 40 {
		nilaiHuruf = "E"
	} else {
		nilaiHuruf = "E"
	}

	fmt.Println("Nilai Anda: ", averageNilai)
	fmt.Println("Nilai Huruf: ", nilaiHuruf)
}