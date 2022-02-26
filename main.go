package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr []int = []int{2, 3, 1, 5}

	fmt.Println(Solution(arr))
}

func Solution(A []int) int {

	sort.Ints(A)
	var res int
	for i := 0; i < len(A)-1; i++ {
		if A[i] != (A[i+1] - 1) {
			res = (A[i] + 1)
		}
	}
	return res
}

//A[0] = 2
//A[1] = 3
//A[2] = 1
//A[3] = 5
