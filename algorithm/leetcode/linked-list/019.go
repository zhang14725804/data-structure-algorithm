

/*
	给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
	给定的 n 保证是有效的。

	思路：双指针（快慢指针）

	（1）建立虚拟头节点（省去判断是否是头节点）
	（2）让某个指针（假设first）向后走n步
	（3）first，second两个指针同时走，相差n步，当first走到结尾时，终止。second走到了倒数第n+1个节点
	（4）倒数n+1个节点直接指向n的下一个节点就好了
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 快慢指针，初始位置都在head
	fast := head
	slow := head
	// 快的指针先走n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// n 等于链表长度
	if fast == nil {
		head = head.Next
		return head
	}
	// 同时向后移动快慢指针
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 删除倒数第n个节点
	slow.Next = slow.Next.Next
	return head
}
