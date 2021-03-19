/*
	填充每个节点的下一个右侧节点指针.完美二叉树
*/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
	方法1：宽度优先遍历（宽搜）,秒啊
*/
func connect1(root *Node) *Node {
	if root == nil {
		return root
	}
	// 缓存root根节点
	source := root
	// 遍历左子树
	for root.Left != nil {
		// 水平遍历
		for p := root; p != nil; p = p.Next {
			// 添加next指针
			p.Left.Next = p.Right
			// (question )链接跨越父节点的两个子节点
			if p.Next != nil {
				p.Right.Next = p.Next.Left
			}
		}
		root = root.Left
	}
	return source
}

/*
	递归
	二叉树的问题难点在于，如何把题目的要求细化成每个节点需要做的事情
*/
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	helper(root.Left, root.Right)
	return root
}
func helper(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	// 前序遍历
	node1.Next = node2
	// 连接两个相同父节点的两个子节点
	helper(node1.Left, node1.Right)
	helper(node2.Left, node2.Right)
	// (question 😅)链接跨越父节点的两个子节点
	helper(node1.Right, node2.Left)
}