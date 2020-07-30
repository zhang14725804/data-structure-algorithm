/*
	填充每个节点的下一个右侧节点指针 II（普通二叉树，有的节点为空）

	两个指针，一个指向头，一个指向尾
*/
type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}
func connect(root *Node) *Node {
	if root==nil{
		return root
	}

	cur:=root
	for cur!=nil{
		head := &Node{}
		tail:=head
		// 不懂，代码有问题！！
		for p:=cur;p!=nil;p=p.Next{
			if p.Left!=nil{
				tail = p.Left
                tail.Next = p.Left
			}
			if p.Right!=nil{
				tail = p.Right
				tail.Next = p.Right
			}

		}
		cur = head.Next
	}
	return root
}