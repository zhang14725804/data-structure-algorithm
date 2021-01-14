/*
	对链表进行插入排序
	插入排序算法：
		插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
		每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
		重复直到所有输入数据插入完为止。

*/
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	prev := head
	node := head.Next
	// 遍历所有节点
	for node != nil {
		// 后一个节点的值小于前一个节点的值
		if node.Val < prev.Val {
			// 从第一个元素开始比较，找到要插入的位置
			cur := dummy
			for cur.Next.Val < node.Val {
				cur = cur.Next
			}
			// 将node节点插入到cur和prev之间
			prev.Next = node.Next
			node.Next = cur.Next
			cur.Next = node
			node = prev.Next
		} else {
			prev = prev.Next
			node = node.Next
		}
	}
	return dummy.Next
}