package main

import "fmt"

func main() {
	// Problem 6
	var n, e int
	fmt.Println("=== Problem 6 - Exponentiation ===")
	fmt.Print("Masukan angka: ")
	fmt.Scanf("%d\n", &n)
	fmt.Print("Masukan pangkat: ")
	fmt.Scanf("%d\n", &e)
	fmt.Printf("%d pangkat %d sama dengan %d", n, e, Eksponensial(n, e))
}
func Eksponensial(base int, pangkat int) int {
	for i := 1; i < pangkat; i++ {
		base *= base
	}
	return base
}
