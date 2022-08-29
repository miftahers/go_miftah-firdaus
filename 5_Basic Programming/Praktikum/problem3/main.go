package main

import "fmt"

func main() {

	// Problem 3
	var (
		f        []int
		bilangan int
	)
	fmt.Println("=== Problem 3 - Faktor Bilangan ===")
	fmt.Print("Masukan bilangan: ")
	fmt.Scanf("%d\n", &bilangan)
	fmt.Printf("Faktor dari %d adalah ", bilangan)
	f = FaktorBilangan(bilangan)
	for i := range f {
		fmt.Printf("%d ", f[i])
	}

}
func FaktorBilangan(n int) (f []int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			f = append(f, i)
		}
	}
	return
}
