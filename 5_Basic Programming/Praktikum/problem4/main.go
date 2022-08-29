package main

import "fmt"

func main() {
	// Problem 4
	var (
		number int
		prima  bool
	)
	fmt.Println("=== Problem 4 - Identifikasi Bilangan Prima ===")
	fmt.Print("Masukan angka: ")
	fmt.Scanf("%d\n", &number)
	prima = BilanganPrima(number)
	if prima {
		fmt.Println("Bilangan Prima")
	} else {
		fmt.Println("Bukan Bilangan Prima")
	}
}

func BilanganPrima(number int) bool {
	var b bool
	b = true
	if number == 0 || number == 1 {
		b = false
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				b = false
				break
			}
		}
	}
	return b
}
