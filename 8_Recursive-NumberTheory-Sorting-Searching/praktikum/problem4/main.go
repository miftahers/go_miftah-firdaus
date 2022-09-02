package main

func MaxSequence(arr []int) int {
	dihitung := len(arr)
	var (
		tmp  int
		tarr []int
	)
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j >= 0; j-- {
			tmp = arr[j] + arr[j-1]
			tarr = append(tarr, tmp)
		}
	}
}
