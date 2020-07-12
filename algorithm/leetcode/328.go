/*
	给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。
	请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

	请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

	应当保持奇数节点和偶数节点的相对顺序。
	链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。

	思路：双指针
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, even := head, head.Next
	// 保留偶数头结点
	evenHead := head.Next
	for even != nil && even.Next != nil {
		// 奇数偶数分别处理
		odd.Next = odd.Next.Next
		odd = odd.Next
		even.Next = even.Next.Next
		even = even.Next
	}
	// 奇数的next指向偶数的头结点
	odd.Next = evenHead
	return head
}