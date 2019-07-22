package main

import "fmt"

func removeOuterParentheses(S string) string {
	var result = make([]byte,0,16)
	var matchCount int
	bs := []byte(S)
	var isOpen bool
	for _,b := range bs {
		if b == 40 && isOpen == false {
			isOpen = true
			continue
		}
		if b == 41 && matchCount == 0 {
			isOpen = false
			continue
		}
		// match count
		if b == 40 && isOpen == true {
			matchCount++
		}
		// isopen   &&     parenthesis not match
		if b == 41 && isOpen == true {
			matchCount--
		}
		result = append(result,b)
	}
	return string(result)
}

func main() {
	s:="()()"
	r:=removeOuterParentheses(s)

	fmt.Println(r)
}
