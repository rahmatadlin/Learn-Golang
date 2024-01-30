package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Bambang"
	names[1] = "Eko"
	names[2] = "Budi"
	// Tidak boleh melebih kapasitas si Array

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// var values = [3]int {
	// 	90,
	// 	80,
	// 	97, // Wajib pakai koma di akhir
	// }

	// Kalau dikosongkan maka defaule values nya adalah 0

	// fmt.Println(values)

	var values = [...]int {
		90,
		98,
		95,
		100,
		104,
	}

	fmt.Println(len(values))

}