package main

import "fmt"

func main() {
	var input string
	fmt.Println("=== Problem 5 - Palindrome ===")
	fmt.Print("Masukan kata: ")
	fmt.Scanf("%s", &input)
	b := Palindrome(input)
	if b {
		fmt.Printf("%s => Palindrome", input)
	} else {
		fmt.Printf("%s => Bukan Palindrome", input)
	}
}
func Palindrome(input string) bool {
	wt := -1
	for i := 0; i < len(input); i++ {
		wt++
	}
	var b bool = true
	wmax := wt
	for wo := 0; wo <= wmax/2; wo++ {
		if input[wo] == input[wt] {
			wt--
		} else {
			b = false
		}
	}
	return b
}
