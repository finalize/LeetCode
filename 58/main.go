// Length of Last Word

package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("The test word is this."))
}

// 0 ms	2.2 MB
func lengthOfLastWord(s string) int {
	c, l := 0, 0
	for _, b := range s {
		if b == 32 {
			c = 0
		} else {
			c++
			l = c
		}
	}
	return l
}
