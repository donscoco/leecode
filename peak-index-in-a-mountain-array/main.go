package main

import "fmt"

//https://leetcode.com/problems/peak-index-in-a-mountain-array

func peakIndexInMountainArray(A []int) int {
	_, s, _ := divideHalf(A, 0, len(A)-1)
	return s
}

func divideHalf(A []int, start, end int) ([]int, int, int) {
	if start == end {
		return A, start, end
	}
	mid := (end-start)/2 + start
	if A[mid] > A[mid+1] {
		return divideHalf(A, start, mid)
	} else {
		return divideHalf(A, mid+1, end)
	}
}

func main() {
	a := []int{3, 4, 5, 1}
	fmt.Println(peakIndexInMountainArray(a))
}
