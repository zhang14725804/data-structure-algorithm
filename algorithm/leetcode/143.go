/*
	给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
	将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

	你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

	方法1：把链表存储到线性表中，然后用双指针依次从头尾取元素即可
	方法2：将链表平均分成两半，将第二个链表逆序，依次连接两个链表（我想到的也是这个方法，不过我是把整个链表反转了😅）
	方法3：递归
*/
func reorderList2(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	// 用快慢指针，寻找链表中间节点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 切分链表
	tail := slow.Next
	slow.Next = nil

	tail = reverseList(tail)
	for tail != nil {
		// 缓存
		tNext := tail.Next
		// 交叉拼接（这里没懂😅）
		tail.Next = head.Next
		head.Next = tail
		head = tail.Next
		// 遍历下个节点
		tail = tNext
	}
}
func reverseList(head *ListNode) *ListNode {
	cur, prev := head, head
	for cur != nil {
		// 缓存next
		next := cur.Next
		// 反转
		cur.Next = prev
		// 移动prev、cur
		prev = cur
		cur = next
	}
	head.Next = nil
	return prev
}

func reorderList3(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	n := 0
	h := head
	// 计算节点个数
	for h != nil {
		h = h.Next
		n++
	}
	dfs(head, n)
}

// todo：递归没看懂😅
func dfs(head *ListNode, n int) *ListNode {
	// 递归出口
	if n == 1 {
		tail := head.Next
		head.Next = nil
		return tail
	}

	if n == 2 {
		tail := head.Next.Next
		head.Next.Next = nil
		return tail
	}

	tail := dfs(head.Next, n-2)
	sub := head.Next
	head.Next = tail
	outTail := tail.Next
	tail.Next = sub
	return outTail
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	list := make([]*ListNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	// 双指针遍历数组，改变指针方向
	start, end := 0, len(list)-1
	for start < end {
		list[start].Next = list[end]
		start++
		if start == end {
			break
		}
		list[end].Next = list[start]
		end--
	}
	list[start].Next = nil
}