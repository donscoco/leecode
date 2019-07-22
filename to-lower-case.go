package main

import "fmt"

func toLowerCase(str string) string {
	bs := []byte(str)
	//    [65,91)    [97,123)    32
	for i,b := range bs {
		if b>=65&&b<91 {
			bs[i]=b+32
		}
	}
	return string(bs)
}


func main() {
	s:="Hello"
	b := []byte(s)
	b[0] = b[0]+32
	fmt.Println(string(b))
	fmt.Println(s)
}
