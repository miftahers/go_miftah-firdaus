package main

import "fmt"

/*
	1. Buat slice lop(list of prime) penampungan daftar prima
	2. if input <= 0 maka return -1 (invalid)
	3. Jadikan kuadrat input sebagai `lim``
	4. buat loop untuk mencari bilangan prima kurang dari `lim` => while i := 2 then i++
	5. jika ditemukan bilangan prima, append slice `lop`
	6. jika `len(lop)`/jumlah dari prima di slice `lop` == input lakukan `break`
	7. jika len`(lop)`` < 1 return -1; else: return lop[number-1]
*/

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

func primeX(number int) int {
	// TODO buat var lop []int untuk menyimpan slice bilangan prima
	var lop []int
	// TODO jika kurang dari nol return -1 yang berarti input harus uint
	if number <= 0 {
		return -1
	} else {
		// TODO buat var lim yang berisi kuadrat dari number + 1 untuk menjadi batas pengecekan nilai int
		lim := number*number + 1
		// TODO loop dari angka 2 sampai dengan lim dan i++ untuk pengecekan nilai i selanjutnya
		for i := 2; i <= lim; i++ {
			// TODO  cek apakah i adalah bilangan prima, jika iya append var lop
			if primeNumber(i) == true {
				lop = append(lop, i)
			}
			//TODO jika panjang lop sebesar nilai inputan maka stop/break looping untuk tidak melakukan perulangan yang tak perlu
			if len(lop) == number {
				break
			}
		}
		// TODO jika panjang lop kurang dari 1 return -1 yang berarti tidak ditemukan bilangan prima
		if len(lop) < 1 {
			return -1
		} else {
			// TODO jika panjang loop lebih dari sama dengan 1 return lop[number-1]
			return lop[number-1]
		}
	}
}
func main() {
	fmt.Println(primeX(1))  // 5
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
