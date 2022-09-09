/*
	Problem 6 - Substitution Cipher

	Implementasikan interface yang terdiri dari metode encode dan decode.
	Algoritma Enkripsi yang digunakan adalah substitusi cipher

	Sample Test Case
	Input:
	[1] Encrypt
	[2] Decrypt
	Choose your menu ? 1
	Input Students Name : rizky
	Output:
	Encode of Students Rizky is irapb
*/

package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	var result []string
	str := strings.ToLower(s.name)

	for i, _ := range str {
		if str[i] <= 109 {
			s := string(110 + (109 - str[i]))
			result = append(result, s)
		} else {
			s := string(109 - (str[i] - 110))
			result = append(result, s)
		}
	}
	nameEncode = strings.Join(result, nameEncode)
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	var result []string
	str := strings.ToLower(s.name)

	for i, _ := range str {
		if str[i] <= 109 {
			s := string(110 + (109 - str[i]))
			result = append(result, s)
		} else {
			s := string(109 - (str[i] - 110))
			result = append(result, s)
		}
	}
	nameDecode = strings.Join(result, nameDecode)
	return nameDecode
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)
	switch {
	case menu == 1:
		fmt.Print("\nInput Students Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode Students Name ", a.name, " is : ", c.Encode())
	case menu == 2:
		fmt.Print("\nInput Students Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode Students Name ", a.name, " is : ", c.Decode())
	default:
		fmt.Println("Wrong input name menu")
	}
}
