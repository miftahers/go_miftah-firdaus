/*
Problem 5 - Students Score

Buat sebuah struct dengan nama Student yang mempunyai properti name dan score dalam bentuk slice kemudian simpan data siswa sebanyak 5 siswa dimasukan maka program menunjukan skor rata-rata, siswa yang memiliki skor minimum dan maksimal?(Implementasikan method)

NOTES: Selesaikan permasalahan ini dengan menggunakan struktur data map
*/
package main

import (
	"fmt"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	sum := 0
	for _, e := range s.score {
		sum += e
	}
	average := sum / len(s.score)
	return float64(average)
}

func (s Student) Min() (min int, name string) {
	var index int
	for i := 0; i < len(s.score); i++ {
		switch {
		case i == 0:
			min, index = s.score[i], i
		case min > s.score[i]:
			min, index = s.score[i], i
		}
	}
	return min, s.name[index]
}

func (s Student) Max() (max int, name string) {
	var index int
	for i := 0; i < len(s.score); i++ {
		switch {
		case i == 0:
			max, index = s.score[i], i
		case max < s.score[i]:
			max, index = s.score[i], i
		}
	}
	return max, s.name[index]
}

func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Input " + string(i) + " Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + string(i) + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Student is ", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Student's is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Student's is : "+nameMin+" (", scoreMin, ")")
}
