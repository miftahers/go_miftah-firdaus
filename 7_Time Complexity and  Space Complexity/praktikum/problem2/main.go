package main

import (
	"fmt"
)

func pow(x, n int) int {
	ans := 1
	for n >= 1 {
		if n%2 == 1 {
			ans *= x
			x *= x
			n = n / 2
		} else {
			ans *= x
			n = n / 2
		}
	}
	return ans
}
func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
