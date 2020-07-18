package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

	层序遍历，以层为单位做dfs。以队列的方式存储每一层
*/
func levelOrder(root *TreeNode) [][]int {
	// 定义答案数组
	var res [][]int
	if root == nil {
		return res
	}
	// 定义队列
	var q Queue
	// 入队
	q.push(*root)
	
	for len(q) > 0 {
		// tips:循环的时候不能直接 i:=0;i<len(q);i++
		length := len(q)
		// 存储每一层元素
		var level []int
		// 遍历当前队列
		for i := 0; i < length; i++ {
			// 出队
			t := q.front()
			// 添加元素
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
		// 将当前层存入res中
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

// 出队（先进先出）并返回出队的元素。指针和地址
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
