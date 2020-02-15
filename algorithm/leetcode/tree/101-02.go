package leetcode

/**
101. Symmetric Tree
判断镜像二叉树问题

迭代（bfs）的方式（不同的遍历方式，然后比较节点即可）
左边：左中右遍历
右边：右中左遍历
*/
func isSymmetric02(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var left Stack
	var right Stack

	l := root.Left
	r := root.Right

	for left.size() > 0 || right.size() > 0 || r != nil || l != nil {
		for r != nil && l != nil {
			left.push(*l)
			l = left.Left
			right.push(*r)
			r = right.Right
		}

		if l == nil || r == nil {
			return false
		}

		l = left.pop()
		r = right.pop()
		if l.Val != r.Val {
			return false
		}

		// 这里注意
		l = l.Right
		r = r.Left
	}
	return true
}

// 利用golang slice实现栈
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack []TreeNode

func (stack *Stack) size() int {
	return len(stack)
}
func (s *Stack) push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) pop() *TreeNode {
	theStack := *s
	if len(theStack) == 0 {
		return nil
	}
	node := theStack[len(theStack)-1]
	*s = theStack[0 : len(theStack)-1]
	return node
}
