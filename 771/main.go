// Jewels and Stones

package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aAbB", "AAAaaaabbbbAAAAddssbbBBBBBB"))
}

// 0 ms	2 MB
func numJewelsInStones(J string, S string) int {
	var c int
	for _, v := range S {
		for _, w := range J {
			if v == w {
				c++
			}
		}
	}

	return c
}
