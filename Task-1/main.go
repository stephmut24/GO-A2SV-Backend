package main

import "fmt"



func SumOfNumber(nums []int)int {
	if len(nums) == 0 {
		return 0
	}

	sum := 0

	for _, value :=range nums {
		sum += value
	}
	return sum

}


func main(){
	fmt.Println(SumOfNumber([]int{1,2,3,4,5}))
	fmt.Println(SumOfNumber([]int{}))

}