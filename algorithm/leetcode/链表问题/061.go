/*
	注意理解，每翻转一次，把最后一个数字往前移动一次，由末尾移到首位
	1. 对进行取模 k%=lstLen
	2. 快慢指针，first指针先从头往后走k步
	3. first和second同时往后走，当first走到尾部的时候，停止
	4. 现在的尾first.Next指向原来的头，原来的头指向新的头second.Next
	5. second截断
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 求出链表长度
	lstLen := 0
	p := head
	for p != nil {
		lstLen++
		p = p.Next
	}
	// 😅
	k %= lstLen
	first := head
	second := head

	for i := k; i > 0; i-- {
		first = first.Next
	}
	// 😅
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	// 现在的尾指向原来的头
	first.Next = head
	// 原来的头指向新的头
	head = second.Next
	// 切断循环
	second.Next = nil
	return head
}
