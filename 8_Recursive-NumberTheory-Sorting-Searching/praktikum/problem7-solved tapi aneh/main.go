package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {

	// 1. cari angka terbesar dari slice deck
	maxDeck, minDeck := FindMax(deck)
	var temp []int

	// 2. Find similar number to maxDeck or minDeck
	for i := 0; i < len(cards); i++ {
		for j := 0; j < 2; j++ {
			if maxDeck == cards[i][j] {
				temp = append(temp, cards[i]...)
			}
		}
	}

	// Find similar number to minDeck
	if len(temp) < 1 {
		for i := 0; i < len(cards); i++ {
			for j := 0; j < 2; j++ {
				if minDeck == cards[i][j] {
					if i == 0 {
						temp = append(temp, cards[i]...)
					}
				}
			}
		}
	}

	if len(temp) < 1 {
		return "tutup kartu"
	} else {
		for i := 0; i < len(temp); i++ {
			for j := 0; j < 2; j++ {
			}
		}
	}
	return temp
}
func main() {
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3})) // [3, 4]
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6})) // [6, 5]
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))         // tutup kartu

}

func FindMax(arr []int) (int, int) {
	var max, min int
	for i, e := range arr {
		switch {
		case i == 0:
			max = e
		case e > max:
			max = e
		case e < min:
			min = e
		}
	}
	return max, min
}

/*
	for i := 0; i < len(cards); i++ {
		for j := 0; j < 2; j++ {
			if deck[0] == cards[i][j] || deck[1] == cards[i][j] {
				temp = append(temp, cards[i])
			}
		}
	}
*/

/*
var res []int
maxTemp, _ := FindMax(temp)
for i := 0; i < len(temp); i += 2 {
	if temp[i] == maxTemp && i != 0 {
		res = append(res, temp[i-1], temp[i])
	} else if temp[i+1] == maxTemp && i != len(temp)-1 {
		res = append(res, temp[i], temp[i+1])
	}
}
*/
