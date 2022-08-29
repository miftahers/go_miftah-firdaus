package main

import (
	"fmt"
	"math"
)

func main() {
	// Problem 1
	var t, r float64
	fmt.Println("=== Problem 1 - Luas Permukaan Tabung ===")
	fmt.Print("Masukan tinggi tabung: ")
	fmt.Scanf("%f\n", &t)
	fmt.Print("Masukan jari-jari alas tabung: ")
	fmt.Scanf("%f\n", &r)
	fmt.Printf("Luas permukaan tabung adalah %.2f satuan\n\n", LuasPermukaanTabung(t, r))
}

func LuasPermukaanTabung(t float64, r float64) (l float64) {
	l = math.Phi * t * r
	return
}
