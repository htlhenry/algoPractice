package main

import "fmt"

// 平均算法复杂度为O(nlogn)
// TODO 算法复杂度推导公式
func quickSort(a []int) []int {
	return quickHelper(0, len(a)-1, a)
}

func quickHelper(left, right int, a []int) []int {
	// 递归代码的模板首先考虑递归终止条件, 这是面试的考点
	if left >= right {
		return a
	}
	// 这个是快排的主要逻辑, 找到一个分区数, 在这个分区数的左边都比它小, 右边都比它大
	// 它的左边的数都比它小不代表已经排好序了, 只代表着分区数的位置已经确定是正确的
	p := quickPartition(left, right, a)
	// 到这里, 这个分区数就已经确定好了自己的位置, 它是不需要在移动的, 并且在它左边的都比它小
	// 所以可以分别再对左边和右边做同样逻辑的排序
	quickHelper(left, p-1, a)
	quickHelper(p+1, right, a)
	return a
}

func quickPartition(left, right int, a []int) int {
	// j代表的是边界数的最终位置, 目前是临时从左开始赋值步进寻找它正确的位置
	i, j := left, left
	// 这里的逻辑就是以right作为边界, 简单的选择分区数的方法, 不一定是最优的分区数选择方法
	for i < right && j < right {
		// 这里就是从左边第一个数开始
		if a[i] < a[right] {
			a[i], a[j] = a[j], a[i]
			j++
		}
		i++
	}
	a[j], a[right] = a[right], a[j]
	return j
}

func main() {
	// 像我这个例子就是个极端情况
	// 第一次寻找分区数, right=4，值为10, 它的位置就是0, 因为它最小
	// 它会遍历到20的位置终止遍历, 与它最终的j交换位置, 其实就是0, 因为它最小
	// 这样第一次遍历的结果就是 [10, 40, 30, 20, 50]
	// 第二次就是只有p+1(p=0)也就是递归a[1:len(a)-1]这个区间, 也是同理开始迭代
	a := []int{50, 40, 30, 20, 10}
	fmt.Printf("quick sort: %d\n", quickSort(a))
}
