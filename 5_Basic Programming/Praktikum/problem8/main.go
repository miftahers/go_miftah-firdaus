package main

import "fmt"

func main() {
	var angka int
	fmt.Println("=== Problem 7 - Play With Arterisk ===")
	fmt.Print("Masukan jumlah angka perkalian: ")
	fmt.Scanf("%d\n", &angka)
	CetakTablePerkalian(angka)
}
func CetakTablePerkalian(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			if j*i < 10 {
				fmt.Printf("   %d", j*i)
			} else if j*i >= 10 && j*i < 100 {
				fmt.Printf("  %d", j*i)
			} else if j*i >= 100 && j*i < 1000 {
				fmt.Printf(" %d", j*i)
			}
		}
		fmt.Println()
	}
}
