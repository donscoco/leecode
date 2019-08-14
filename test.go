package main

import (
	"fmt"
	"time"
)

func map_test() {
	maps := make(map[int][][]byte)
	maps[1]=[][]byte{}
	maps[2]=append(maps[2],[]byte{})
	maps[0]=nil
	for k,v := range maps {
		fmt.Printf("  k: %d  ,  v:  %v  ",k,v)
	}
}

func string_equal_test() {
	s1 := "0"
	s2 := "0"
	fmt.Println(s1 == s2)
}

func map_interface() {
	m := make(map[int]interface{})

	array := [][][]byte{}
	array = append(array,[][]byte{[]byte{1}})
	maps  := make(map[string][][]byte)
	//maps["key1"] = [][]byte{[]byte{2}}

	m[1] = array
	m[2] = maps

	for i,v:= range m[1].([][][]byte){
		fmt.Printf(" index: %d , value: %v ",i,v)
	}
	for k,v:= range m[2].(map[string][][]byte){
		fmt.Printf(" key: %s , value: %v ",k,v)
	}
}
func switch_case_muti(){
	array := []int{1,2,3,4,5}
	for _,v := range array {
		switch v {
		case 1,3:
			fmt.Println("A")
		case 2,4:
			fmt.Println("B")
		}
	}
}

func select_return(){
	dataList := make(chan []byte)
	stop := make(chan bool)
	var count int
	var isStop bool
	go func(){
		for {
			time.Sleep(3*time.Second)
			dataList<-[]byte{123}
			time.Sleep(3*time.Second)
			stop<-true
		}
	}()
	for {
		select{
		case data := <-dataList:
			fmt.Println(data)
			return
		case isStop := <-stop:
			fmt.Println(isStop)
		}
		fmt.Println(count)
		count++
	}
	fmt.Println(isStop)
}

func main(){
	// map_test
	//string_equal_test()
	//map_interface()
	//switch_case_muti()
	select_return()
}