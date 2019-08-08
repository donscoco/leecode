package main

// https://leetcode.com/problems/squares-of-a-sorted-array/

func sortedSquares(A []int) []int {
	result := make([]int,len(A))
	// find the middle number
	var l,m,r int
	if len(A) == 1 {
		result[0] = squares(A,0)
		return result
	}
	if len(A) == 2 {
		if squares(A,0) < squares(A,1) {
			result[0]=squares(A,0)
			result[1]=squares(A,1)
		}else{
			result[0]=squares(A,1)
			result[1]=squares(A,0)
		}
		return result
	}
	l,m,r = localMiddle(A)
	result[0] = squares(A,m)
	//
	i := 1
	for {
		for l >= 0 && r <= len(A)-1 && squares(A,l) <= squares(A,r) {
			result[i] = squares(A,l)
			l--
			i++
		}
		for l >= 0 && r <= len(A)-1 && squares(A,l) > squares(A,r) {
			result[i] = squares(A,r)
			r++
			i++
		}
		for l == -1 && r < len(A) {
			result[i] = squares(A,r)
			r++
			i++
		}
		for r == len(A) && l > -1 {
			result[i] = squares(A,l)
			l--
			i++
		}
		if i == len(A) {
			break
		}
	}

	return result
}
func squares (A []int,i int) int {
	return A[i]*A[i]
}
func localMiddle (A []int) (left,middle,right int) {
	for i := 1; i < len(A); i++ {
		if squares(A,middle) > squares(A,i){
			middle = i
		}
	}
	return middle-1,middle,middle+1
}
func main() {
	//array := []int{5}
	//array := []int{-5,-1,-1,1,5}
	//array := []int{-1,2,2}
	//array := []int{-3,-3,-2,1}
	//fmt.Println(sortedSquares(array1))

	//fmt.Println(localMiddle(array))
}
