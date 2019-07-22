package main

import (
	"fmt"
)

//https://leetcode.com/problems/unique-morse-code-words/

var morse = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
var morseMap = make(map[string]byte)
var byteMap = make(map[byte]string)

func uniqueMorseRepresentations(words []string) int {
	var result int
	// init
	var byteMap = make(map[byte][]byte)
	var existMap = make(map[string]bool)
	for i,v :=range morse {
		byteMap[byte(i+97)]=[]byte(v)
	}
	// process
	for _,word := range words {
		var morseOfWords []byte
		for _,b := range []byte(word){
			ms := byteMap[b]
			morseOfWords = append(morseOfWords,ms...)
		}
		key := string(morseOfWords)
		if !existMap[key] {
			result++
			existMap[key]=true
		}
	}
	return result
}

func initMap(){
	for i,v :=range morse {
		morseMap[v]=byte(i+97)
		byteMap[byte(i+97)]=v
	}
}

func getByte(morse string) byte{
	return morseMap[morse]
}
func getMorse(key byte) string{
	return byteMap[key]
}


func main() {
	initMap()
	fmt.Println(getMorse([]byte("i")[0]))
	fmt.Println(string(getByte("-.")))

	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}
