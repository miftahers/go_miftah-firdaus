package main

import "fmt"

func PairSum(arr []int, target int) []int {
	var res []int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if sum := arr[i] + arr[j]; sum == target && arr[i] != arr[j] {
				res = append(res, i)
			}
		}
	}
	return res
}
func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) //[1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  //[0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   //[2,3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   //[1,2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    //[0,1]
}
