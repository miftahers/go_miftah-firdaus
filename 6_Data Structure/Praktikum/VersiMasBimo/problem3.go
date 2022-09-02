package problem3

import "fmt"

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) //[1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  //[0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   //[2,3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   //[1,2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    //[0,1]
}

func PairSum(arr []int, target int) []int {
	// TODO temp var target value penjumlahan
	temp := make(map[int]int)

	// TODO loop over slice
	for i, number := range arr {
		// TODO target == number
		if key, exist := temp[target-number]; exist {
			return []int{key, i}
		}
		temp[number] = i
	}
	return nil
}
