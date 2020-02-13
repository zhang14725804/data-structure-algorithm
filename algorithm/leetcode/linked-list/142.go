package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	142. Linked List Cycle II

	这个有难度，思路太强
	快慢指针，开始的时候一个走两步一个走一步，相遇之后，慢的指针从开头继续走，另一个再相遇点开始走，每次都走一步，一定会在b点相遇
*/
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	for slow != nil {
		fast = fast.Next
		slow = slow.Next

		if slow != nil {
			slow = slow.Next
		} else {
			break
		}

		// 第一次相遇
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}
