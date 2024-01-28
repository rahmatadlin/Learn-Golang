package main

import "fmt"

func main () {
	// Saat membuat variable wajib menyebutkan tipe data variable tersebut

	// var name string
	// var name = "Rahmat Adlin"
	name := "Rahmat Adlin" // Hanya deklarasi pertama saja. tidak boleh membuat untuk kedua kalinya, coba yg kedua pasti error
	fmt.Println(name)

	name = "Rahmat"
	fmt.Println(name)

	// Berikut untuk multiple variable
	// Kalau varibale hanya di declare saja tapi tidak digunakan maka akan error
	var (
		firstName = "Rahmat"
		lastName = "Adlin"
	)


	fmt.Println(firstName, lastName)
	fmt.Println(lastName)

}