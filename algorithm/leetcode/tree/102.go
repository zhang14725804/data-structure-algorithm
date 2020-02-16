package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
	102. Binary Tree Level Order Traversal
	层序遍历，以层为单位做dfs。以队列的方式存储每一层，和101-02栈的思路一样
*/
func levelOrder(root *TreeNode) [][]int {
    
}

// 利用slice实现队列
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