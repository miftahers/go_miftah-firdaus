package main

import (
	"fmt"
	"sort"
)

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})             // 3
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 14000, 5000, 3000}) // 4
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})          // 4
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 3000})                  // 1
	MaximumBuyProduct(0, []int{10000, 30000})                               // 0
}

func MaximumBuyProduct(money int, productPrice []int) {
	var count, sum int = 0, 0

	// urut productPrice
	sort.Ints(productPrice)
	for i := 0; i < len(productPrice); i++ {
		// Cek apakah bisa beli i barang
		if sum+productPrice[i] <= money {
			sum += productPrice[i]
			// tambah var penghitung
			count++
		}
	}
	fmt.Println(count)
}
