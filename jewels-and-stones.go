package main

func numJewelsInStones(J string, S string) int {
	//convert
	var result int
	js := []byte(J)
	ss := []byte(S)
	pocket := make(map[byte]bool,50)
	for _,j := range js {
		pocket[j]=true
	}
	for _,s := range ss {
		if pocket[s]==true {
			result++
		}
	}
	return result
}

func main() {
	
}
