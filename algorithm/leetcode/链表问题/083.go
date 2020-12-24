/*
	给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
	思路：快慢指针
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	// 移动快慢指针
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	// 断开后续重复数据
	slow.Next = nil
	return head
}