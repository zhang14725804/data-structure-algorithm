/*
	给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
*/

/*
	方法1：快慢指针 😅😅😅😅😅😅
	依次遍历即可 😮😮😮
	（1）建立虚拟头节点（省去判断是否是头节点）
	（2）让某个指针（假设first）向后走n步
	（3）first，second两个指针同时走，相差n步，当first走到结尾时，终止。second走到了倒数第n+1个节点
	（4）倒数n+1个节点直接指向n的下一个节点就好了
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 快慢指针，初始位置都在head
	fast, slow := head, head
	// （1）快的指针先走n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// （2）n 等于链表长度
	if fast == nil {
		return head.Next
	}
	// （3）同时向后移动快慢指针，注意这里要判断【fast.Next】是否走到结尾位置
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// （4）删除倒数第n个节点
	slow.Next = slow.Next.Next
	return head
}

/*
	方法2：两次遍历
	删除一个结点，无非是遍历链表找到那个结点前边的结点，然后改变下指向就好了。
	（1）但由于它是链表，它的长度我们并不知道，我们得先遍历一遍得到它的长度，
	（2）之后用长度减去 n 就是要删除的结点的位置，然后遍历到结点的前一个位置就好了。
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lLen := 0
	// 😅 这里要用虚拟头节点进行遍历
	p := head
	// 获取链表长度
	for p != nil {
		p = p.Next
		lLen++
	}
	// 长度等于1的情况 ，再删除一个结点就为 null 了
	if lLen == 1 {
		return nil
	}

	rmNodeIndex := lLen - n
	// 如果删除的是头结点
	if rmNodeIndex == 0 {
		return head.Next
	}

	// question 删除链表的第n个节点(12.23号,字节面试被问到)
	// 😅 这里要用虚拟头节点进行遍历
	p = head
	// 😅找到被删除结点的前一个结点
	for i := 0; i < rmNodeIndex-1; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	return head
}

/*
	方法3：
	第一次遍历链表确定长度的时候，顺便把每个结点存到数组里，这样找结点的时候就不需要再遍历一次了
	空间换时间？？？
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 统计链表长度
	lLen := 0
	listArray := make([]*ListNode, 0)
	p := head
	// 将链表塞进slice
	for p != nil {
		lLen++
		listArray = append(listArray, p)
		p = p.Next
	}

	// 总共有一个节点的情况
	if lLen == 1 {
		return nil
	}

	// 删除第一个节点的情况
	rmi := lLen - n
	if rmi == 0 {
		return head.Next
	}

	rmNodePrev := listArray[rmi-1]
	rmNodePrev.Next = rmNodePrev.Next.Next
	return head

}