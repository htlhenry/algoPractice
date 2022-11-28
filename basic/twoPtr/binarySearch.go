package main

import "fmt"

// BinarySearch
// 双端闭合的搜索区间, [left, right]
func BinarySearch(a []int, target int) int {
	left, right := 0, len(a)-1
	for left <= right {
		mid := left + (right-left)/2
		if a[mid] == target {
			return mid
		} else if a[mid] > target {
			right = mid - 1
		} else if a[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

// LeftBound
// 搜索左边界, 搜索等于target的最左边的索引
// 搜索区间, [left, right)
func LeftBound(a []int, target int) int {
	left, right := 0, len(a)
	for left < right { // 左开有闭, 终止条件 left == right, right = len(a)
		mid := left + (right-left)/2
		if a[mid] == target {
			right = mid // 找到之后继续往左靠拢, 为什么不是mid-1, 因为右开, mid-1就会漏掉
		} else if a[mid] < target {
			left = mid + 1 // 为什么是mid+1, 左闭, 此时mid已经搜索过了, 所以要跳过
		} else if a[mid] > target {
			right = mid // 同 == target
		}
	}
	return left
}

func main() {
	fmt.Printf("binary search: %d \n", BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4))

	fmt.Printf("left bound: %d\n", LeftBound([]int{1, 2, 2, 2, 2, 3, 4, 5}, 2))
}
