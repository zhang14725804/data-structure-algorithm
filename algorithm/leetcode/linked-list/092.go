/*
	92.Reverse Linked List II

	链表反转某一段

	思路：
	（1）保存m的前一个元素，保存n的下一个元素
	（2）反转m-n之间的元素，
	（3）修改m、n指针指向
*/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	// 第一个节点有可能变，所以先声明一个虚拟头结点
	newHead := &ListNode{}
	newHead.Next = head

	// a 是m的前一个元素
	a := newHead
	// 保存 n元素
	d := newHead

	for i := 0; i < m-1; i++ {
		a = a.Next
	}
	for i := 0; i < n; i++ {
		d = d.Next
	}

	// b 表示m元素
	b := a.Next
	// 保存n的下一个元素
	c := d.Next

	// m到n之间元素指针翻转
	prev := b
	cur := b.Next
	for cur != c {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	b.Next = c
	a.Next = d

	return newHead.Next
}
