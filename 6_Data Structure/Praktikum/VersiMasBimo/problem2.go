package problem1

import (
	"fmt"
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
	// TODO Loop over string == []byte
	// TODO temp var isi nya key angka elemen jumlahnya
	temp := make(map[int]int)

	for _, char := range angka {
		str := string(char)
		tempInt, _ := strconv.Atoi(str)
		temp[tempInt]++
	}

	// TODO bikin var return
	result := make([]int, 0)
	// TODO seleksi yang dia di dalam temp cuman 1 aja
	for key, element := range temp {
		if element == 1 {
			result = append(result, key)
		}
	}
	return result
}
