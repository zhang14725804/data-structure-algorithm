/*
	给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
	层序遍历的基础上，反转
*/

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var q Queue
	// 指针OK
	q.push(*root)

	for len(q) > 0 {
		// tips:循环的时候不能直接 i:=0;i<len(q);i++
		length := len(q)
		var level []int
		for i := 0; i < length; i++ {
			t := q.front()
			level = append(level, t.Val)

			if t.Left != nil {
				q.push(*t.Left)
			}
			if t.Right != nil {
				q.push(*t.Right)
			}
		}
		res = append(res, level)
	}
	// 反转结果
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
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