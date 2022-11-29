package main

type LinkNode struct {
	Val  int
	Next *LinkNode
}

// MergeTwoSortedLinkList
// 合并两个有序链表
// 链表的技巧, 虚拟头节点
func MergeTwoSortedLinkList(p1, p2 *LinkNode) *LinkNode {
	dummyHead := &LinkNode{Val: -1}
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
	return dummyHead.Next
}
