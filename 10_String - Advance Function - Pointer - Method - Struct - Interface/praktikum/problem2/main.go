/*
Problem 2 - Caesar Cipher

menggeser huruf
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Caesar(3, "abc"))                           // def
	fmt.Println(Caesar(2, "alta"))                          // cnvc
	fmt.Println(Caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(Caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(Caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}

func Caesar(offset int, input string) string {

	str := strings.ToLower(input)
	tmp := offset % 26
	var result []string

	for i, _ := range str {
		if str[i]+byte(tmp) <= 122 {
			s := string((str[i]) + byte(tmp))
			result = append(result, s)
		} else {
			a := 122 - str[i]
			b := byte(tmp) - a
			s := string(b + 96)
			result = append(result, s)
		}
	}
	return strings.Join(result, "")
}
