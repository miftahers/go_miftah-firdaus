/*
Problem 4 - Min and Max Using Pointer

Tulis program di Golang untuk menemukan nilai maksimum serta minimum di antara 6 angka inputan.
Gunakan multiple return fungsi, pointer untuk referencing maupun dereferencing!

Sample Test Case
Input:
1
2
3
9
7
8
Output:
9 is the maximum number
1 is the minimum number
*/
package main

import "fmt"

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai Min: ", min)
	fmt.Println("Nilai max: ", max)
}

func getMinMax(numbers ...*int) (min, max int) {
	for i, e := range numbers {
		switch {
		case i == 0:
			max, min = *e, *e
		case *e > max:
			max = *e
		case *e < min:
			min = *e
		}
	}
	return max, min
}
