/*
	给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
	将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
	你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
	0105 懵逼 0106 懵逼
	question 😅😅😅😅😅😅
*/

/*
	方法1：把链表存储到线性表中，然后用双指针依次从头尾取元素即可（说的轻松 😅）
*/
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	list := make([]*ListNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	// 双指针遍历数组，依次改变指针方向（注意指针指向 😅😅😅）
	start, end := 0, len(list)-1
	for start < end {
		// （1）改变 start 指针
		list[start].Next = list[end]
		start++
		if start == end {
			break
		}
		// （2）改变 end 指针
		list[end].Next = list[start]
		end--
	}
	list[start].Next = nil
}

/*
	方法2：将链表平均分成两半，将第二个链表逆序，依次连接两个链表
	（说的轻松 😅），第六步很复杂
*/
func reorderList(head *ListNode) {
	// （1） 边界条件
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	// （2） 虚拟头结点
	dummy := new(ListNode)
	dummy.Next = head
	// （3） 快慢指针，找到中间节点
	fast, slow := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// （4） 后半部分
	tail := slow.Next
	// 切断前后两段之间的链接
	slow.Next = nil
	// （5）反转后半部分
	tail = reverseList(tail)

	// （6） 拼接前后两部分（顺序不能乱 最绕的 😅😅😅😅😅😅😅😅）
	for tail != nil {
		tNext := tail.Next    // 缓存 tail.Next
		tail.Next = head.Next // 链接 tail.Next和head.Next
		head.Next = tail      // 链接tail和head
		head = tail.Next      // head = head.Next
		tail = tNext          // tail = tail.Next
	}
}

/*
	方法3：递归（ 更难懂 😅😅😅）
*/
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

