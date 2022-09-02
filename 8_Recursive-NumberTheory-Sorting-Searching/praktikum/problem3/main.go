/*
Segi empat berisi bilangan prima:

	Yang diperlukan:
		- high * wide = max => max isi slice prima
		- slice[0] > start
		- Nested for a.k.a O(n^2) untuk membuat segi empat
*/

package main

import "fmt"

func primaSegiEmpat(wide, high, start int) {
	// to get max len slice
	max := high * wide
	// to get primesSlice dan total
	primes, total := primesSliceWithTotal(max, start)
	// to print segi empat primes
	temp := 0
	for i := 0; i < wide; i++ {
		for j := 0; j < high; j++ {
			fmt.Printf("%d ", primes[temp])
			temp++
		}
		fmt.Println()
	}
	// to print total
	fmt.Println(total)
}

func main() {
	primaSegiEmpat(2, 3, 13)
	/*
		17 19
		23 29
		31 37
		156
	*/
	primaSegiEmpat(5, 2, 1)
	/*
		2 3 5 7 11
		13 17 19 23 29
		129
	*/
}

// to cek prima
func primeNumber(number int) bool {
	if number == 1 {
		return false
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

// To get primesSlice with customized start and total sum of slice value
func primesSliceWithTotal(number, start int) ([]int, int) {
	var lop []int
	total := 0
	if number <= 0 {
		return nil, 0
	} else {
		for i := start + 1; i > 0; i++ {
			// cek prima
			if primeNumber(i) == true {
				lop = append(lop, i)
				total += i
			}
			if len(lop) == number {
				break
			}
		}
		if len(lop) < 1 {
			return nil, 0
		} else {
			return lop, total
		}
	}
}
