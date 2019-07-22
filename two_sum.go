package main

func twoSum(nums []int, target int) []int {
	listMap := make(map[int][]int,1000)
	for i,value := range nums{
		list := listMap[value]
		listMap[value]=append(list,i)
	}
	for i,value := range nums{
		next:= target - value
		list:= listMap[next]
		if len(list)>0&&list[0]!=i{
			return []int{i,list[0]}
		}
		if len(list)>1&&list[0]==i{
			return []int{i,list[1]}
		}
		continue
	}
	return make([]int,2)
}

func main() {
	
}
