package main

import (
	"fmt"
)

func main() {
	// 只出现一次的数字
	// var arr = []int{1, 2, 2, 3, 3}
	// num := onceNum(arr)
	// fmt.Println(num)

	//符号处理
	// str := "{{}]"
	// str1 := "[]"
	// fmt.Println(dealStr(str))
	// fmt.Println(dealStr(str1))

	// 公共前缀
	// strArr := []string{"flower", "flow", "flight"}
	// fmt.Println(longestCommonPrefix(strArr))

	// 最大数+1
	numArr := []int{4, 3, 2, 1}
	fmt.Println(plusOne(numArr))

}

// 最大数+1
func plusOne(digits []int) []int {
	n := len(digits)
	// 从最后一位（最低位）开始遍历，处理进位
	for i := n - 1; i >= 0; i-- {
		// 当前位加 1
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}

	return append([]int{1}, digits...)
}

// 最长公共前缀
func longestCommonPrefix(strArr []string) string {
	if len(strArr) == 0 {
		return ""
	}

	first := strArr[0]
	for i := 0; i < len(first); i++ {
		char := first[i]
		for j := 0; j < len(strArr); j++ {
			if i > len(strArr[j]) || strArr[j][i] != char {
				return first[:i]
			}
		}
	}
	return first
}

// 符号处理
func dealStr(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	mp := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []rune{}

	for _, v := range str {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != mp[v] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
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
