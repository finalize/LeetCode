package main

import "fmt"

func main() {
	arr := []int{5, 1, 5, 2, 5, 3, 5, 4}
	fmt.Println(repeatedNTimes(arr))
}

func repeatedNTimes(A []int) int {
	var res int
	c := make(map[int]int)
	for _, v := range A {
		if _, ok := c[v]; ok {
			res = v
			break
		}
		c[v] = 1
	}
	return res
}
