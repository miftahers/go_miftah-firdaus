/*
	Problem 1 - Compare String

	You are given two string A and B. Determine common substring between them.

	Sample Test Case
	Input: A = "AKA" B = "AKASHI"
	Output: AKA

	Sample Test Case
	Input: A = "KANGOORO" B = "KANG"
	Output: "KANG"
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}

func Compare(a, b string) string {
	if strings.Contains(a, b) {
		return b
	}
	if strings.Contains(b, a) {
		return a
	}
	return "Tidak ada kesamaan"
}
