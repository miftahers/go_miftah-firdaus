package main

import "fmt"

func main() {

	// Problem 2
	var studentScore int
	var studentName string
	fmt.Println("=== Problem 2 - Grade Nilai ===")
	fmt.Print("Masukan nama mahasiswa: ")
	fmt.Scanf("%s\n", &studentName)
	fmt.Print("Masukan nilai: ")
	fmt.Scanf("%d\n", &studentScore)
	fmt.Printf("%s = %d(%s)\n", studentName, studentScore, GradeNilai(studentScore))

}
func GradeNilai(n int) (grade string) {
	if n >= 80 && n <= 100 {
		grade = "Nilai A"
	} else if n >= 65 && n <= 79 {
		grade = "Nilai B"
	} else if n >= 50 && n <= 64 {
		grade = "Nilai C"
	} else if n >= 35 && n <= 49 {
		grade = "Nilai D"
	} else if n >= 0 && n <= 34 {
		grade = "Nilai E"
	} else {
		grade = "Tidak Valid"
	}
	return
}
