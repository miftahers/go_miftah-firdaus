/*
find the most appear item
1. Buat map[string]int dimana key nya adalah value dari arr[] input dan int nya adalah jumlah key tersebut pada array
2. loop in range arr[]
*/
package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func mostAppearItem(items []string) []pair {
	// Buat map dan masukan isi items []string ke map tersebut
	tstring := make(map[string]int)
	for _, item := range items {
		_, ok := tstring[item]
		if ok {
			tstring[item]++
		} else {
			tstring[item]++
		}
	}

	// buat keys sebagai penampungan jumlah dari map
	keys := make([]string, 0, len(tstring))

	for key := range tstring {
		keys = append(keys, key)
	}
	// Sort keys
	sort.SliceStable(keys, func(i, j int) bool {
		return tstring[keys[i]] < tstring[keys[j]]
	})

	// Uji coba slice of struct
	var (
		p  pair
		sp []pair
	)
	for _, k := range keys {
		p.name = k
		p.count = tstring[k]
		sp = append(sp, p)
	}
	return sp
}

func main() {
	fmt.Println(mostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})) // Golang->1 Ruby->2 js->3
	fmt.Println(mostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})) // c->1 D->2 B->3 A->4
	fmt.Println(mostAppearItem([]string{"Football", "Basketball", "Tennis"}))               // Football->1 Basketball->1 Tennis->1
}

/*

	// UJI COBA TANPA SERCING

	// Buat map dan masukan isi items []string ke map tersebut
	tstring := make(map[string]int)
	for _, item := range items {
		_, ok := tstring[item]
		if ok {
			tstring[item]++
		} else {
			tstring[item]++
		}
	}

	keys := make([]string, 0, len(tstring))

	for key := range tstring {
			keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool{
			return tstring[keys[i]] < tstring[keys[j]]
	})

	for _, k := range keys{
			fmt.Printf("%v->%v ",k, tstring[k])
	}
	return nil
*/
