/*
	Challange - Letter frequency
	Hitung frekuensi huruf dalam teks menggunakan perhitungan paralel(bersamaan)

	Sample Test Case
	Input:
	Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua

	Output:
	e : 1
	i : 1
	o : 1
	t : 1
	...
*/

package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var alphabet []string
	str := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	for i := 0; i < len(str); i++ {
		alphabet = append(alphabet, string(str[i]))
	}

	m := make(map[string]int)
	mMutex := sync.RWMutex{}

	cfh := func() {
		mMutex.Lock()
		for i := 0; i < len(alphabet)/2; i++ {
			m[alphabet[i]]++
		}
		mMutex.Unlock()
		wg.Done()
	}

	clh := func() {
		mMutex.Lock()
		for i := len(alphabet) / 2; i < len(alphabet); i++ {
			m[alphabet[i]]++
		}
		mMutex.Unlock()
		wg.Done()
	}

	wg.Add(1)
	go cfh()

	wg.Add(1)
	go clh()

	wg.Wait()

	for k, e := range m {
		fmt.Printf(`%v: %v`+"\n", k, e)
	}
}
