package main

import (
	"fmt"
)

// O(sqrt(n))
func primeNumber(number int) bool {
	// jika satu bukan prima
	if number == 1 {
		return false
	}
	// loop dari 2 ke sqrt(x)
	for i := 2; i*i < number; i++ {
		// check apakah i membagi x tidak ada sisa
		if number%i == 0 {
			// jika iya maka ada faktor pembagi x diantara 2 dan sqrt(x)
			// oleh karena itu nilai disini bukan prima
			return false
		}
	}
	// jika tidak menemukan faktor lain di perulangan dengan sqrt(x) maka number adalah prima
	return true
}
func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
	fmt.Println(primeNumber(1500450271)) // true
}
