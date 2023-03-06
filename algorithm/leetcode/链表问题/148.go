/*
	方法1:归并排序（这种方法优美）；
	归并排序（MERGE-SORT）是利用归并的思想实现的排序方法，该算法采用经典的分治（divide-and-conquer）策略
	分治法将问题：
		分(divide)成一些小的问题然后递归求解
		治(conquer)的阶段则将分的阶段得到的各答案"修补"在一起，即分而治之
	方法2（蠢办法）：将链表节点放入数组，将数组排序，之后再链接各个排序后的节点
*/
func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	// 😅 base case，递归出口
	if head == nil || head.Next == nil {
		return head
	}
	// 😅😅😅 question 为什么这里需要虚拟头节点。不用会 out of memory
	dummy := new(ListNode)
	dummy.Next = head
	fast, slow := dummy, dummy

	// 😅😅 快慢指针找到中点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tail := slow.Next
	// 😅😅😅 切断两段之间的链接
	slow.Next = nil
	// 😅 递归前半部分head
	head = mergeSort(head)
	// 😅 递归后半部分tail
	tail = mergeSort(tail)
	// 合并前后两部分
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
			// 各自向前走
			tail = tail.Next
			head1 = head1.Next
		} else {
			tail.Next = head2
			// 各自向前走
			tail = tail.Next
			head2 = head2.Next
		}
	}
	//  😅 拼接剩余的部分
	if head1 != nil {
		tail.Next = head1
	}
	if head2 != nil {
		tail.Next = head2
	}
	return dummy.Next
}
