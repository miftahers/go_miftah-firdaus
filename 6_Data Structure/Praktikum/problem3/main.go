package main

import "fmt"

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) //[1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  //[0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   //[2,3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   //[1,2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    //[0,1]
}

func PairSum(arr []int, target int) []int {
	var res []int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			// TODO bandingkan penjumlahan arr[i] dan arr[j] dengan target
			if sum := arr[i] + arr[j]; sum == target {
				// jika i == j jangan append karena index yang sama
				if i != j {
					res = append(res, i)
				}
			}
		}
	}
	return res
}
