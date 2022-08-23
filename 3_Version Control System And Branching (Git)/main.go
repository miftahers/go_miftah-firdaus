package main

import "fmt"

func main() {
	fmt.Println("Hello Aku")
	fmt.Println("seluruh dunia")
	// stash
	fmt.Println("Dan Dunia")

	// Menambahkan looping membuat segitiga siku-siku
	for i := 1; i < 10; i++ {
		for j := 1; j < i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
