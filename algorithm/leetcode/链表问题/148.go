/*
	给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
	进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

	😅😅😅😅😅😅
	question
	归并排序（MERGE-SORT）是利用归并的思想实现的排序方法，该算法采用经典的分治（divide-and-conquer）策略
	（分治法将问题分(divide)成一些小的问题然后递归求解，而治(conquer)的阶段则将分的阶段得到的各答案"修补"在一起，即分而治之)
*/
func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 😅😅😅 为什么这里需要虚拟头节点（当节点个数是偶数的时候，让 slow 刚好指向前边一半节点的最后一个节点）
	dummy := new(ListNode)
	dummy.Next = head
	fast, slow := dummy, dummy
	// 快慢指针找到中点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tail := slow.Next
	// 😅😅😅 切断两段之间的链接
	slow.Next = nil
	// 前半部分
	head = mergeSort(head)
	// 后半部分
	tail = mergeSort(tail)
	return merge(head, tail)
}

func merge(head1, head2 *ListNode) *ListNode {
	// 😅😅😅 又是虚拟头节点
	dummy := new(ListNode)
	// 😅😅😅 这个节点是专门用来遍历的
	tail := dummy
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			tail.Next = head1
			tail = tail.Next
			head1 = head1.Next
		} else {
			tail.Next = head2
			tail = tail.Next
			head2 = head2.Next
		}
	}
	if head1 != nil {
		tail.Next = head1
	}
	if head2 != nil {
		tail.Next = head2
	}
	return dummy.Next
}
