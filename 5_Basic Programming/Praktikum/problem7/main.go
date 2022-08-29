package main

import "fmt"

func main() {
	var bintang int
	fmt.Println("=== Problem 7 - Play With Arterisk ===")
	fmt.Print("Masukan jumlah bintang terbanyak: ")
	fmt.Scanf("%d\n", &bintang)
	PlayWithAsterisk(bintang)
}
func PlayWithAsterisk(n int) {
	nmax := n
	fmt.Println()
	for row := 1; row <= n; row++ {
		for space := 1; space <= nmax; space++ {
			fmt.Print(" ")
			continue
		}
		for star := 1; star <= row; star++ {
			fmt.Print("* ")
		}
		nmax--
		fmt.Println()
	}
	fmt.Println()
}
