package main

import (
	"fmt"
)

func main() {
	num := 10
	pointerAdd10(&num)
	fmt.Println(num)

	arr := []int{1, 2, 3}
	splice2(&arr)
	fmt.Println(arr)
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func splice2(arr *[]int) {
	if arr == nil || len(*arr) == 0 {
		return
	}

	for i, v := range *arr {
		(*arr)[i] = v * 2
	}
}

// 值+10
func pointerAdd10(numPointer *int) {
	*numPointer += 10
}
