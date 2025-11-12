package main

import (
	"fmt"
)

func main() {
	var arr = []int{1, 2, 2, 3, 3}
	num := onceNum(arr)
	fmt.Println(num)
}

// 只出现一次的数字
func onceNum(arr []int) int {

	numMap := make(map[int]int)
	for num := range arr {
		value, exist := numMap[arr[num]]
		if exist {
			value++
			numMap[arr[num]] = value
		} else {
			numMap[arr[num]] = 1
		}
	}

	for key, value := range numMap {
		if value == 1 {
			fmt.Println("只重复一次的元素为：", key)
			return key

		}
	}
	return 0
}
