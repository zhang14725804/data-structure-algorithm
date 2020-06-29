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
	// res:=make([][]int)
	// 声明二维数组（上面的错了）
	var res [][]int
	if root == nil {
		return res
	}

	var q Queue
	q.push(*root)
	for len(q) > 0 {
		length := len(q)
		//level:=make([]int)
		// 声明二维数组（上面的错了）
		var level []int
		for i := 0; i < length; i++ {
			t := q.front()
			level = append(level, t.Val)
			if t.Left != nil {
				// 注意是指针
				q.push(*t.Left)
			}
			if t.Right != nil {
				// 注意是指针
				q.push(*t.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

// 利用slice实现队列
type Queue []TreeNode

// 入队
func (s *Queue) push(node TreeNode) {
	*s = append(*s, node)
}

// 出队
func (s *Queue) front() *TreeNode {
	theStack := *s
	node := &TreeNode{}
	if len(theStack) == 0 {
		return node
	}
	node = &theStack[0]
	*s = theStack[1:len(theStack)]
	return node
}
