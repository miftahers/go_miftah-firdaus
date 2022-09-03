/*
find the most appear item
1. Buat map[string]int dimana key nya adalah value dari arr[] input dan int nya adalah jumlah key tersebut pada array
2. loop in range arr[]
*/
package main

import "fmt"

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
	fmt.Println(tstring)
	//puts tstring elements into tmp slice
	var tmp []int
	for _, e := range tstring {
		tmp = append(tmp, e)
	}
	// Sort tmp
	for i, e := range tmp {
		if i != len(tmp)-1 {
			if e > tmp[i+1] {
				temp := tmp[i+1]
				tmp[i+1] = tmp[i]
				tmp[i] = temp
			}
		}
	}

	// print sorted
	for _, se := range tmp {
		for k, me := range tstring {
			if se == me {
				fmt.Printf("%v->%v ", k, me)
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(mostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})) // Golang->1 Ruby->2 js->3
	fmt.Println(mostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})) // c->1 D->2 B->3 A->4
	fmt.Println(mostAppearItem([]string{"Football", "Basketball", "Tennis"}))               // Football->1 Basketball->1 Tennis->1
}
