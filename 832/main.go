package main

func main() {
	arr := [][]int{[]int{1, 0, 0, 0}, []int{1, 1, 0, 1}, []int{1, 1, 1, 0}}
	flipAndInvertImage(arr)
}

// myself
func flipAndInvertImage(A [][]int) [][]int {
	var res [][]int
	for _, a := range A {
		var arr []int
		for i := len(a)/2 - 1; i >= 0; i-- {
			opp := len(a) - 1 - i
			a[i], a[opp] = a[opp], a[i]
		}
		for _, v := range a {
			if 0 < v {
				v--
			} else {
				v++
			}
			arr = append(arr, v)
		}
		res = append(res, arr)
	}
	return res
}

// simplest
// func flipAndInvertImage(A [][]int) [][]int {
// 	for _, s := range A {
// 		for i := 0; i < (len(A)+1)/2; i++ {
// 			s[i], s[len(A)-i-1] = s[len(A)-i-1]^1, s[i]^1
// 		}
// 	}
// 	return A
// }
