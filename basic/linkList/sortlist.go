package main

// leetcode 148 sort list

// 自顶向下的解法, 时间复杂度O(n *logn), 空间复杂度O(N*LogN)

// 自下向上的解法, 时间复杂度O(n * Logn), 空间复杂度On
func mergeTwoListFromBottom(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var length int
	for tmp := head; tmp != nil; tmp = tmp.Next {
		length++
	}
	dummyHead := &ListNode{Next: head}
	for subLength := 1; subLength < length; subLength <<= 1 { // 按照1, 2, 4(每次分割长度*2), 按照subLength分离然后两两子链表合并排序, 用迭代替换递归占用的空间
		pre, cur := dummyHead, dummyHead.Next // pre 表示链表的头节点 cur表示当前的头节点
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil // 分离head1
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil // 分离head2
			}
			cur = next
			pre.Next = MergeTwoSortedListNode(head1, head2) // 合并排序两个head
			for pre.Next != nil {
				pre = pre.Next
			}
		}
	}
	return dummyHead.Next
}
