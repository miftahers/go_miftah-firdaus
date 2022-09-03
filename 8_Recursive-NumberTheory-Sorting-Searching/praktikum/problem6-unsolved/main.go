/*
2. loop productPrice => if productPrice[i] < money, append temp[]int
3. sort temp dan jumlahkan dari barang termurah sampai termahal
4. break jika total > money return i+1
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})             // 3
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 14000, 5000, 3000}) // 4
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 3000})
	MaximumBuyProduct(0, []int{10000, 30000})
}

func MaximumBuyProduct(money int, productPrice []int) {
	var tempAsc, tempDesc []int
	var count int
	for _, e := range productPrice {
		if e < money {
			tempAsc = append(tempAsc, e)
			tempDesc = append(tempDesc, e)
		}
	}
	if len(tempAsc) < 1 {
		fmt.Println(count)
	} else {
		sort.Ints(tempAsc)
		sort.Slice(tempDesc, func(i, j int) bool {
			return tempDesc[i] > tempDesc[j]
		})
		total := make([]int, 0)
		for _, ae := range tempAsc {
			for _, de := range tempDesc {
				if FindMax(total) <= money {
					for i := 0; i < len(total); i++ {

					}
					count++
				} else {
					break
				}
			}
		}
		fmt.Println(count - 1)
	}
}

func FindMax(arr []int) int {
	var max int
	for i, e := range arr {
		switch {
		case i == 0:
			max = e
		case e > max:
			max = e
		}
	}
	return max
}

// 1 max
