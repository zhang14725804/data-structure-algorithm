/*
	填充每个节点的下一个右侧节点指针.完美二叉树

	思路：宽度优先遍历（宽搜）
*/
type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}

func connect(root *Node) *Node {
	if root == nil{
		return root
	}
	source:=root
	// 有些绕
	for root.Left != nil{
		for p:=root;p!=nil;p=p.Next{
			// left的next指向right
			p.Left.Next = p.Right
			if p.Next != nil{
				// 左节点的right的next指向右节点的left
				p.Right.Next = p.Next.Left
			}
		}
		// 递归下一个左节点
		root=root.Left
	}
	return source
}