package main

import "testing"

func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		numJewelsInStones("aAbB", "AAAaaaabbbbAAAAAaaaabbbbAAAAddssbbBBBBBBAAAaaaabbbbAAAAddssbbBBBBBBAAAaaaabbbbAAAAddssbbBBBBBBAAAaaaabbbbAAAAddssbbBBBBBBAAAaaaabbbbAAAAddssbbBBBBBBAAAaaaabbbbAAAAddssbbBBBBBBAAAaaaabbbbAAAAddssbbBBBBBBAAddssbbBBBBBB")
	}
}

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
