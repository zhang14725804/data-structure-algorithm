type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	判断镜像二叉树问题

	迭代（bfs）的方式（自行维护一个栈）

	左边：左中右遍历
	右边：右中左遍历

	todos：：指针*和地址&
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var left Stack
	var right Stack

	l := root.Left
	r := root.Right

	for len(left) > 0 || len(right) > 0 || r != nil || l != nil {

		for r != nil && l != nil {
			// 左子树的左子树 和 右子树的右子树
			left.push(*l)
			l = l.Left
			right.push(*r)
			r = r.Right
		}

		// 左右节点一个为空一个不为空
		if l != nil || r != nil {
			return false
		}

		l = left.pop()
		r = right.pop()
		// 左右子节点val不同
		if l.Val != r.Val {
			return false
		}

		// 左子树的右子树 和 右子树的左子树
		l = l.Right
		r = r.Left
	}
	return true
}

// 利用slice实现栈
type Stack []TreeNode

func (s *Stack) push(node TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) pop() *TreeNode {
	theStack := *s
	node := &TreeNode{}
	if len(theStack) == 0 {
		return node
	}
	node = &theStack[len(theStack)-1]
	*s = theStack[0 : len(theStack)-1]
	return node
}
