/*
	Leetcode061 Rotate List。注意理解，每翻转一次，把最后一个数字往前移动一次，由末尾移到首位
	1. k%=n
	2. first指针从头往后走k步
	3.first和second同时往后走，当first走到尾部的时候，停止

	todos：：我还是不懂快慢指针
*/
func Leetcode061(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 求出链表长度
	n := 0
	p := head
	for p != nil {
		n++
		p = p.Next
	}

	k %= n
	first := head
	second := head

	for i := k; i > 0; i-- {
		first = first.Next
	}
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	// 现在的尾指向nil，现在的头指向原来的头
	first.Next = head
	head = second.Next
	second.Next = nil
	return head
}
