package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6,3]
	fmt.Println(munculSekali("12345"))      // [1, 2, 3, 4, 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8, 7, 2, 5, 4]
}
func munculSekali(angka string) []int {
	var (
		slice []int
		num   int
		res   []int
	)
	// TODO buat slice int isinya string angka O(n)
	for i := 0; i < len(angka); i++ {
		num, _ = strconv.Atoi(string(angka[i]))
		slice = append(slice, num)
	}
	// TODO Urut isi slice angka
	sort.Ints(slice)

	// TODO perulangan perbandingan antar index O(n)
	var n int = len(slice)
	for i := 0; i < n-1; i++ {
		if i > 0 {
			// TODO bandingin angka index setelah awal dan sebelum akhir
			if slice[i] != slice[i+1] && slice[i] != slice[i-1] {
				res = append(res, slice[i])
			}
		} else if i == 0 {
			// TODO bandingin angka index awal sama index setelah awal
			if slice[0] != slice[1] {
				res = append(res, slice[0])
			}
		}
	}
	// TODO bandingin angka index terakhir sama index sebelum terakhir
	if slice[n-1] != slice[n-2] {
		res = append(res, slice[n-1])
	}
	// Kembalikan nilai hasil (var res)
	return res
}
