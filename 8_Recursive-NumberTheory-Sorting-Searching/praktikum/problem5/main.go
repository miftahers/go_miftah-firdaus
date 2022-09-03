package main

import (
	"fmt"
	"strconv"
)

func FindMinMax(arr []int) string {
	var min, max, imax, imin int
	for i, e := range arr {
		switch {
		case i == 0:
			max, min, imax, imin = e, e, i, i
		case e > max:
			max, imax = e, i
		case e < min:
			min, imin = e, i
		}
	}
	strmin, strmax, strimin, strimax := strconv.Itoa(min), strconv.Itoa(max), strconv.Itoa(imin), strconv.Itoa(imax)
	return "min: " + strmin + " index: " + strimin + " max: " + strmax + " index: " + strimax
}
func main() {
	fmt.Println(FindMinMax([]int{5, 7, 4, -2, -1, 8}))
	//min: -2 index: 3 max: 8 index: 5
	fmt.Println(FindMinMax([]int{2, -5, -4, 22, 7, 7}))
	//min: -5 index: 1 max: 22 index: 3
	fmt.Println(FindMinMax([]int{4, 3, 9, 4, -21, 7}))
	//min: -21 index: 4 max: 9 index: 2
	fmt.Println(FindMinMax([]int{-1, 5, 6, 4, 2, 18}))
	// min: -1 index: 0 max: 18 index: 5
	fmt.Println(FindMinMax([]int{-2, 5, -7, 4, 7, -20}))
	// min: -20 index: 5 max: 7 index: 4
}
