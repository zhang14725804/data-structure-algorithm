/*
	给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
	如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
	您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
	0109 自己写的一塌糊涂
*/
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	// 虚拟头节点
	dummy := &ListNode{}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		// 当前和
		curSum := carry
		// 遍历两个链表，相加节点值
		if l1 != nil {
			curSum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			curSum += l2.Val
			l2 = l2.Next
		}
		// 进位
		carry = curSum / 10
		// curSum % 10当前位的数值
		// 😅😅😅 这里需要初始化cur.Next节点，不可以直接写成 cur.Next
		cur.Next = &ListNode{Val: curSum % 10}
		// cur前进一步
		cur = cur.Next
	}
	return dummy.Next
}

/**
	0109 自己写的时候遇到的错误点
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    dummy:=&ListNode{}
    pointer:=dummy
    flag:=0
    for l1!=nil && l2!=nil{
		// (1)  l1.Val+l2.Val+flag 忘记加 flag进位
		// 😅😅😅 这里需要 初始化pointer.Next节点，不可以直接写成 pointer.Val 
        pointer.Next = &ListNode{Val:(l1.Val+l2.Val+flag) % 10}
        flag = (l1.Val+l2.Val+flag) / 10
        pointer = pointer.Next
        l1 = l1.Next
        l2 = l2.Next
    }
	// (2) for 循环写成了 if
    for l1 != nil {
        pointer.Next = &ListNode{Val:(l1.Val+flag) % 10}
        flag = (l1.Val+flag) / 10
        pointer = pointer.Next
        l1 = l1.Next
    }
    for l2 != nil {
        pointer.Next = &ListNode{Val:(l2.Val+flag) % 10}
        flag = (l2.Val+flag) / 10
        pointer = pointer.Next
        l2 = l2.Next
    }
	// （3）最后的进位忘记
    if flag>0{
        pointer.Next = &ListNode{Val:flag}
        pointer = pointer.Next
    }
    return dummy.Next
}