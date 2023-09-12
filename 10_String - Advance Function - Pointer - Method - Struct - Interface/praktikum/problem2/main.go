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
	fmt.Println(CaesarIgnoreSpace(3, "abc"))                           // def
	fmt.Println(CaesarIgnoreSpace(2, "alta"))                          // cnvc
	fmt.Println(CaesarIgnoreSpace(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(CaesarIgnoreSpace(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(CaesarIgnoreSpace(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
	fmt.Println(CaesarCipherRefactored(7, "afila ansori"))
	fmt.Println((CaesarCipher(-7, "hmpsh huzvyp")))
}

func CaesarIgnoreSpace(offset int, input string) string {

	str := strings.ToLower(input)
	tmp := offset % 26
	var result []string

	for i := range str {
		if str[i] == byte(' ') {
			continue
		}
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

func CaesarCipherRefactored(shift int, input string) string {
	inputLowerCase := strings.ToLower(input)
	shiftAmount := shift % 26
	var encryptedChars []string

	for i := range inputLowerCase {
		char := inputLowerCase[i]

		if char+byte(shiftAmount) <= 'z' {
			shiftedChar := string(char + byte(shiftAmount))
			encryptedChars = append(encryptedChars, shiftedChar)
		} else {
			remainingShift := 'z' - char
			wrappedShift := byte(shiftAmount) - remainingShift
			wrappedChar := string('a' + wrappedShift - 1)
			encryptedChars = append(encryptedChars, wrappedChar)
		}
	}

	return strings.Join(encryptedChars, "")
}

func CaesarCipher(shift int, input string) string {
	shiftAmount := shift % 26
	var encryptedChars []string

	for i := range input {
		char := input[i]

		if 'a' <= char && char <= 'z' {
			if char+byte(shiftAmount) <= 'z' {
				shiftedChar := string(char + byte(shiftAmount))
				encryptedChars = append(encryptedChars, shiftedChar)
			} else {
				remainingShift := 'z' - char
				wrappedShift := byte(shiftAmount) - remainingShift
				wrappedChar := string('a' + wrappedShift - 1)
				encryptedChars = append(encryptedChars, wrappedChar)
			}
		} else if 'A' <= char && char <= 'Z' {
			if char+byte(shiftAmount) <= 'Z' {
				shiftedChar := string(char + byte(shiftAmount))
				encryptedChars = append(encryptedChars, shiftedChar)
			} else {
				remainingShift := 'Z' - char
				wrappedShift := byte(shiftAmount) - remainingShift
				wrappedChar := string('A' + wrappedShift - 1)
				encryptedChars = append(encryptedChars, wrappedChar)
			}
		} else {
			// Preserve non-alphabet characters
			encryptedChars = append(encryptedChars, string(char))
		}
	}

	return strings.Join(encryptedChars, "")
}
