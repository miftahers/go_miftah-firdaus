package problem1

import "fmt"

func ArrayMerge(arrA, arrB []string) []string {
	if len(arrA) < 1 && len(arrB) < 1 {
		return nil
	}

	// TODO bikin temp var buat semua nama, kalau ada duplikat diitung 1 aja
	tempName := make(map[string]int)
	var res []string

	// TODO simpan semua nama dari arrA
	for _, name := range arrA {
		_, exist := tempName[name]
		if !exist {
			tempName[name] = 1
		}
	}

	// TODO simpan semua nama dari arrB
	for _, name := range arrB {
		_, exist := tempName[name]
		if !exist {
			tempName[name] = 1
		}
	}

	for key, _ := range tempName {
		res = append(res, key)
	}

	return res
}
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
