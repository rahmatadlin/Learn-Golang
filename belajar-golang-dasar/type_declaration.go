package main

import "fmt"

func main() {

	type NoKTP string

	var ktpRahmat NoKTP = "123456789"

	var contoh string = "22222222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpRahmat)
	fmt.Println(contohKTP)


}