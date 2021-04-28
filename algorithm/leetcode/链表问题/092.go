/*
	链表反转某一段
*/

/*
	方法1：迭代实现
	（1）保存m的前一个元素，保存n的下一个元素
	（2）反转m-n之间的元素，
	（3）修改m、n指针指向
*/
func reverseBetween1(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	dummy := &ListNode{} // 第一个头节点有可能变
	dummy.Next = head    // 为什么需要这一步 😅😅😅

	mPrev := dummy // m的前一个节点
	nNode := dummy // n节点
	// 😅😅😅找到m上一个节点
	for i := 0; i < m-1; i++ {
		mPrev = mPrev.Next
	}
	// 😅😅😅找到n下一个节点
	for i := 0; i < n; i++ {
		nNode = nNode.Next
	}

	mNode := mPrev.Next // 😅😅😅 m节点
	nNext := nNode.Next // 😅😅😅 n节点的下一个节点

	// 反转m到n之间的元素
	prev := mNode
	cur := mNode
	for cur != nNext {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	// 😅😅😅 careful 改变反转链表的前后指针
	mNode.Next = nNext
	mPrev.Next = nNode
	return dummy.Next
}

/*
	方法2：递归实现
	将链表反转某一段，改为反转链表的前n个节点
*/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// 😅 base case，反转前n个元素
	if m == 1 {
		return reverseN(head, n)
	}
	// 😅 前进到反转的起点触发 base case
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

/*
	反转链表的前n个节点
*/
var last *ListNode // 记录后续的节点

func reverseN(head *ListNode, n int) *ListNode {
	// 😅 base case
	if n == 1 {
		// 😅😅😅 缓存第n+1个节点的指针
		last = head.Next
		return head
	}
	// 以head.Next为起点，反转前n-1个节点
	tail := reverseN(head.Next, n-1)
	head.Next.Next = head
	// 链接反转之后的节点和后面的节点 😅😅😅
	head.Next = last
	// 返回tail 😅😅😅
	return tail
}
