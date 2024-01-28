package main

import "fmt"

func main () {
	// Tidak error apabila hanya di declare saja namun tidak boleh diubah lagi
	const firstName string = "Rahmat"
	const lastName = "Adlin"

	// Maka akan error apabila di re asign
	// Bisa juga multiple variable
	firstName = "Tidak bisa diubah"
	lastName = "Tidak bisa diubah"

	fmt.Println(firstName, lastName)

}