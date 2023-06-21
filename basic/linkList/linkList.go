package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoSortedListNode
// 合并两个有序链表
// 链表的技巧, 虚拟头节点
func MergeTwoSortedListNode(p1, p2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1}
	p := dummyHead
	l1, l2 := p1, p2
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	// 根据上面的循环终止条件, l1或者l2存在以下两种情况:
	// 1: l1或l2为nil
	// 2: l1和l2都为nil
	// 所以这里判断一下如果有不为nil的直接把他赋值到当前node的Next
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return dummyHead.Next
}

func main() {
	p1 := &ListNode{}
	//dp1 := p1
	//for dp1 != nil {
	//    fmt.Printf("dp1 val: %d ", dp1.Val)
	//    dp1 = dp1.Next
	//}

	p2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}
	dp2 := p2
	for dp2 != nil {
		fmt.Printf("dp2 val: %d", dp2.Val)
		dp2 = dp2.Next
	}
	head := MergeTwoSortedListNode(p1, p2)
	for head.Next != nil {
		fmt.Printf("val: %d, ", head.Val)
		head = head.Next
	}

}
