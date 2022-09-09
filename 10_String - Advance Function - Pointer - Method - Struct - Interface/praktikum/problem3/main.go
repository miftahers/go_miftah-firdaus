/*
Problem 3 - Swap Two Number Using Pointer

Fungsi swap number adalah fungsi yang bertugas untuk menukar dua nilai dalam sebuah variabel.
Contoh variabel a memiliki nilai 10, variabel b memiliki nilai 20. Setelah ditukar, a memiliki nilai 20 dan b memiliki nilai 10. Buatkan sebuah fungsi tersebut dengan menggunakan pointer!

Sample Test Case
Input: a = 10; b = 20
Output: a = 20; b = 10
*/
package main

import "fmt"

func main() {
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
