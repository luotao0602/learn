package main

import (
	"fmt"
	"sort"
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
	// numArr := []int{4, 3, 2, 1}
	// fmt.Println(plusOne(numArr))

	//6、删除有序数组中的重复项
	//  nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//  size := removeDuplicates(nums)
	//  fmt.Println(size)

	//合并
	// arr := [][]int{
	// 	{4, 5},
	// 	{1, 4},
	// }
	// merged := merge(arr)
	// fmt.Println(merged)

	// 两数之和
	target := 6
	nums := []int{3, 2, 4}
	res := sum(nums, target)
	fmt.Println(res)

}

// 两数之和
func sum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))

	for i, v := range nums {
		mp[v] = i + 1
	}
	fmt.Println(mp)
	arr := []int{}
	for i, v := range nums {
		a := mp[target-v]
		if a != 0 && i != a-1 {
			arr = append(arr, i)
			arr = append(arr, a-1)
			break
		}
	}
	return arr
}

// 合并
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	// 先按区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果切片，放入第一个区间
	merged := [][]int{intervals[0]}

	for _, current := range intervals[1:] {
		// 获取结果中最后一个区间
		last := merged[len(merged)-1]
		// 如果当前区间的起始位置小于等于结果中最后区间的结束位置，说明有重叠
		if current[0] <= last[1] {
			// 合并区间：更新结果中最后区间的结束位置为两者较大值
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 没有重叠直接添加进新切片中
			merged = append(merged, current)
		}
	}

	return merged
}

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	//不重复元素的位置
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
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
