package main

import "fmt"

/*
归并排序: 原理是将一个数组不断拆分为原来的1/2,拆分到只剩下一个元素的数组返回, 回溯
*/
func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	mid := len(a) / 2
	left := mergeSort(a[:mid])
	right := mergeSort(a[mid:])
	return merge(left, right)
}

func merge(a []int, b []int) []int {
	var res []int
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		res = append(res, a[i])
	}
	for ; j < len(b); j++ {
		res = append(res, b[j])
	}
	return res
}

func main() {
	a := []int{50, 40, 30, 20, 10}
	fmt.Printf("merge sort: %d\n", mergeSort(a))
}
