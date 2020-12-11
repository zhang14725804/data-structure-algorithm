/*
	148. Sort List

	Sort a linked list in O(n log n) time using constant space complexity.
*/
func sortList(head *ListNode) *ListNode {
	// 获取链表长度
	n := 0
	p := head
	for p != nil {
		p = p.Next
		n++
	}
	// 第一个节点有可能变，所以先声明一个虚拟头结点
	newHead := &ListNode{}
	newHead.Next = head

	for i := 1; i < n; i *= 2 {
		cur := newHead
		for j := 0; j+i < n; j += i * 2 {
			left := cur.Next
			right := cur.Next
			for k := 0; k < i; k++ {
				right = right.Next
			}
			l := 0
			r := 0
			for l < i && r < i && right != nil {
				if left.Val <= right.Val {
					cur.Next = left
					cur = left
					left = left.Next
					l++
				} else {
					cur.Next = right
					cur = right
					right = right.Next
					r++
				}
			}
			for l < i {
				cur.Next = left
				cur = left
				left = left.Next
				l++
			}
			for r < i && right != nil {
				cur.Next = right
				cur = right
				right = right.Next
				r++
			}
			cur.Next = right
		}
	}
	return newHead.Next
}
