package main

import "fmt"

func main() {
	arr := []int{-100, -3, -1, 0, 1, 123, 143435, 53453463}
	fmt.Println(sortedSquares(arr))
}

func sortedSquares(A []int) []int {
	length := len(A)
	res := make([]int, length)
	i, j := 0, length-1

	for z := length - 1; z >= 0; z-- {
		sq1, sq2 := A[i]*A[i], A[j]*A[j]

		if sq1 > sq2 {
			res[z] = sq1
			i++
		} else {
			res[z] = sq2
			j--
		}
	}

	return res
}
