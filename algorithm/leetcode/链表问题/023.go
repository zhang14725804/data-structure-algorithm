/*
	合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

	方法二：优先队列（todo：又是优先队列）
	0103
*/

/*
	方法1：两两合并
	TODO: 是否可以改进成多协程的，两个两个合并之后再两个两个合并
*/
func mergeKLists(lists []*ListNode) *ListNode {
	// []和[[]] 分两种情况，不能投篮
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 0 {
		return nil
	}
	// 两两合并
	head := mergeTwoLists(lists[0], lists[1])
	for i := 2; i < len(lists); i++ {
		head = mergeTwoLists(head, lists[i])
	}
	return head
}

// 注意： 参数是*ListNode类型，而不是[]*ListNode
func mergeTwoList(list1,list2 *ListNode) *ListNode{
	// 需要虚拟头结点
    head:=&ListNode{}
    dummy:=head
    for list1!=nil && list2!=nil{
        if list1.Val<list2.Val{
            head.Next = list1
            list1=list1.Next
        }else{
            head.Next = list2
            list2=list2.Next
        }
        head = head.Next
    }

    if list1!=nil{
        head.Next = list1
    }
    if list2!=nil{
        head.Next = list2
	}
	// 最后返回【dummy.Next】而不是dummy
    return dummy.Next
}