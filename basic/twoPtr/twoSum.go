package main

import "fmt"

// TwoSum
// leetCode 167
// 两数之和可以使用双指针的前提是数组已排序
func TwoSum(a []int, target int) []int {
	var res []int
	left, right := 0, len(a)-1
	for left < right {
		sum := a[left] + a[right]
		if sum == target {
			res = append(res, left, right)
			break
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return res
}

func main() {
	idx := TwoSum([]int{2, 7, 10, 12, 15}, 19)
	fmt.Printf("two sum idx: %d\n", idx)
}
