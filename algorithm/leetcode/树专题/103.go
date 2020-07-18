/*
	给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

	思路：层序遍历之后，按照层级反转答案即可
*/
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil{
		return res
	}
	// 
	var q Queue
	q.push(*root)
	cnt:=0
	// 
	for len(q) > 0{
		// 
		length := len(q)
		var level []int
		cnt++
		
		for i := 0; i < length; i++{
			// 
			t:=q.front()
			level = append(level,t.Val)
			
			// 
			if t.Left !=nil{
				q.push(*t.Left)
			}
			if t.Right !=nil{
				q.push(*t.Right)
			}
		}
		// 偶数行反转数组
        if cnt % 2 == 0{
            reverse(level)
        }
		res = append(res,level)
	}
	return res
}

func reverse(res []int){
	// 反转数组
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
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