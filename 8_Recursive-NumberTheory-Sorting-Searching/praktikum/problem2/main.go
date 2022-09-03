package main

import "fmt"

func fibonacci(number int) int {
	// Jika number sama dengan 0 kembalikan nol
	if number == 0 {
		return 0
		// Jika number sama dengan 1 kembalikan 1
	} else if number == 1 {
		return 1
		// Jika number bukan 0 atau 1 return fungsi fibonacci(number-1) + fibonacci(number-2)
		// Nantinya akan melakukan pengecekan berulang dan menjumlahkan 2 nilai sebelum number
	} else {
		return fibonacci(number-1) + fibonacci(number-2)
	}
}
func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}

/*

number := 9

fibbonacci {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	} else {
		return fibonacci(number-1) + fibonacci(number-2)
	}
}

fibonacci(number-1) => number = 8
fibonacci(number-2) => number = 7
	fibonacci(number-1) => number = 6
	fibonacci(number-2) => number = 5
		fibonacci(number-1) => number = 4
		fibonacci(number-2) => number = 3
			fibonacci(number-1) => number = 2
			fibonacci(number-2) => number = 1 => return 1

*/
