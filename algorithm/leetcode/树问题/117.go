/*
	给定一个二叉树
	填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
*/
func connect(root *Node) *Node {
	var queue []*Node
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		cLen := len(queue)
		// 缓存前一个节点 😅😅😅
		var preNode, cnode *Node
		for i := 0; i < cLen; i++ {
			if i == 0 {
				cnode = queue[0]
				queue = queue[1:]
				// preNode首次赋值
				preNode = cnode
			} else {
				cnode = queue[0]
				queue = queue[1:]
				// 增加next指针
				preNode.Next = cnode
				// preNode前进一步
				preNode = cnode
			}
			// 遍历下一层
			if cnode.Left != nil {
				queue = append(queue, cnode.Left)
			}
			if cnode.Right != nil {
				queue = append(queue, cnode.Right)
			}
		}
		preNode.Next = nil
	}
	return root
}