package main

import "fmt"

// Init Struct
type User struct {
	ID int
	Nama string
	email string
	status string
}

func main(){

	// Cara Seperti ini boleh di isi secara acak
	user := User{}
	user.ID = 1
	user.Nama = "Ilham Budiawan"
	user.email = "budiawanilham04@gmail.com"
	user.status = "aktif"

	user2 := User{}
	user2.ID = 2
	user2.Nama = "Khoirunnisa"
	user2.email = "khoirunnisa@gmail.com"
	user2.status = "aktif"

	// sedangkan yang ini harus di isi secara berurutan
	user3 := User{
		ID: 2,
		Nama: "Bela Nabila",
		email: "belanabila@gmail.com",
		status: "aktif",
	}

	// Print hasil
	fmt.Println(user)
	fmt.Println(user2)
	fmt.Println(user3)
}