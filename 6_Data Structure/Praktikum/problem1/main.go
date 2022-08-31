package main

import "fmt"

func main() {
	// test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimatsu", "alisa", "law"}))
	// ["Alisa", "yoshimitsu", "devil jin", "law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}

func ArrayMerge(arrayA, arrayB []string) (result []string) {
	// 1. gabungkan 2 array/slice jadi 1 slice
	// 2. buat map[stirng]bool, (ingat! default bool adalah false) buat variabel sementara untuk keep bool dari map hasil
	// 3. lakukan cek jika false maka set bool ke true (untuk tidak meng-append string selanjutnya yang mirip) kemdian slice res di append

	// Merge Slice
	var temp []string
	temp = append(temp, arrayA...)
	temp = append(temp, arrayB...)

	// Remove Duplicate
	inResult := make(map[string]bool)
	var res []string
	for _, str := range temp {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			res = append(res, str)
		}
	}
	return res
}
