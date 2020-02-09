package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	92.Reverse Linked List II

	翻转某一段
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
	// d 表示n那个元素
	d := newHead

	for i := 0; i < m-1; i++ {
		a = a.Next
	}
	for i := 0; i < n; i++ {
		d = d.Next
	}

	// b 表示m那个元素
	b := a.Next
	// c 是n的下一个元素
	c := d.Next

	// m到n之间元素指针翻转
	p := b
	q := b.Next
	for q != c {
		o := q.Next
		q.Next = p
		p = q
		q = o
	}
	b.Next = c
	a.Next = d

	return newHead.Next
}
