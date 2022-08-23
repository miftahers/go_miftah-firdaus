package main

import "fmt"

func main() {
	fmt.Println("Hello Aku")
	fmt.Println("seluruh dunia")
	// stash
	fmt.Println("Dan Dunia")

	// Membuat fitur perkalian dari input
	var a1 int
	var a2 int
	fmt.Println("========= Perkalian ==========")
	fmt.Println("masukan angka pertama: ")
	fmt.Scanln(&a1)
	fmt.Println("masukan angka kedua: ")
	fmt.Scanln(&a2)
	fmt.Println("Hasilnya adalah", a1*a2)

	// Menambahkan looping membuat segitiga siku-siku
	for i := 1; i < 10; i++ {
		for j := 1; j < i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
