// Max sum nilai di array
package main

import "fmt"

func MaxSeq(arr []int) int {
	max := arr[0]
	acc := 0

	for _, e := range arr {
		acc += e
		if acc > max {
			max = acc
		}
		if acc < 0 {
			acc = 0
		}
	}
	return max
}

func main() {
	fmt.Println(MaxSeq([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSeq([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSeq([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSeq([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSeq([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
}
