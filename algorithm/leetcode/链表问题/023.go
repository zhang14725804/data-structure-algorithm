/*
	方法1：两两合并
*/
func mergeKLists(lists []*ListNode) *ListNode {
	// []和[[]] 分两种情况，不能偷懒
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

/*
	方法二：优先队列（todo：又是优先队列）
	0103 😅😅😅😅
*/