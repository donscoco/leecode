package main


import (
	"fmt"
	_ "net/http/pprof"
	)

//https://leetcode.com/problems/n-repeated-element-in-size-2n-array/

func repeatedNTimes(A []int) int {

	exist := make(map[int]bool)
	scope := len(A)/2+1+1
	for i := 0;i < scope;i++ {
		if !exist[A[i]] {
			exist[A[i]] = true
		}else{
			return A[i]
		}
	}
	return -1
}

func main() {
	A := []int{1,2,3,3}
	result := repeatedNTimes(A)
	fmt.Println(result)
}
