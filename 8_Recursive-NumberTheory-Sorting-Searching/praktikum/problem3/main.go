/*
Segi empat berisi bilangan prima:

	Yang diperlukan:
		- high * wide = max => max isi slice prima
		- slice[0] > start
		- Nested loop O(n^2) untuk membuat segi empat
*/

package main

import "fmt"

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

func primaSegiEmpat(wide, high, start int) {
	// TODO buat max len(slice)
	max := high * wide
	// TODO dapatkan nilai prima dan total dengan panggil fungsi slicePrimaDanTotalSum
	primes, total := slicePrimaDanTotalSum(max, start)
	// TODO nested loop untuk print slice prima dengan bentuk segi empat
	temp := 0
	for i := 0; i < wide; i++ {
		for j := 0; j < high; j++ {
			fmt.Printf("%d ", primes[temp])
			temp++
		}
		fmt.Println()
	}
	// TODO print total
	fmt.Println(total)
}

// TODO buat fungsi cek prima atau bukan
func primeNumber(number int) bool {
	// TODO cek number jika == 1 return false
	if number == 1 {
		return false
	}
	// TODO loop angka dari 2 dengan batasan loop adalah number/2 karena maksimal hasil dari nilai yang bisa dibagi adalah sama dengan 2, i++
	for i := 2; i < number/2; i++ {
		// Cek apakah dibagi nilai i bisa habis, jika iya return false
		if number%i == 0 {
			return false
		}
	}
	// TODO return true jika di cek tidak ditemukan bilangan yang bisa membagi nya
	return true
}

// TODO buat fungsi untuk dapetin slice isinya bil. Prima dan total jumlah semuanya
func slicePrimaDanTotalSum(number, start int) ([]int, int) {
	var lop []int
	total := 0
	if number <= 0 {
		return nil, 0
	} else {
		for i := start + 1; i > 0; i++ {
			// TODO call fungsi cek prima, jika outputnya true maka append lop dan total += i
			if primeNumber(i) == true {
				lop = append(lop, i)
				total += i
			}
			// TODO jika panjang lop == number maka stop loop supaya tidak loop yang tak diperlukan
			if len(lop) == number {
				break
			}
		}
		// TODO jika lop kurang dari 1 return nil dan 0
		if len(lop) < 1 {
			return nil, 0
			// TODO jika lop lebih dari sama dengan 1 return lop dan total
		} else {
			return lop, total
		}
	}
}
