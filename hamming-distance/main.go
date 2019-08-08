package main

//https://leetcode.com/problems/hamming-distance

func hammingDistance(x int, y int) int {
	var result int
	for x != 0 || y != 0 {
		if x%2 != y%2 {
			result++
		}
		x = x >> 1
		y = y >> 1
	}
	for x != 0 && y == 0 {
		if x%2 == 1 {
			result++
		}
		x = x >> 1
	}
	for x == 0 && y != 0 {
		if y%2 == 1 {
			result++
		}
		y = y >> 1
	}
	return result
}

func main() {
	hammingDistance(1,4)
}
