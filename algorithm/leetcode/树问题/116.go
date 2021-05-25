/*
	给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。
	填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
	type Node struct {
		Val   int
		Left  *Node
		Right *Node
		Next  *Node
	}
*/

/*
	方法1：借助【队列】实现层序遍历
*/
func connect(root *Node) *Node {
	var queue []*Node
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		cLen := len(queue)
		// 缓存前一个节点 😅😅😅
		var preNode *Node
		// 当前节点
		var cnode *Node
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

/*
	方法1：宽度优先遍历（宽搜）, 😅😅😅😅😅😅
*/
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	// 缓存root根节点
	source := root
	// 😅遍历左子树
	for root.Left != nil {
		// 😅水平遍历
		for p := root; p != nil; p = p.Next {
			// 添加next指针
			p.Left.Next = p.Right
			// 😅 链接跨越父节点的两个子节点
			if p.Next != nil {
				p.Right.Next = p.Next.Left
			}
		}
		// 😅遍历下一层
		root = root.Left
	}
	return source
}

/*
	方法2：递归实现
	三种情况 😅😅😅：
		（1）链接当前左右节点
		（2）链接两个父节点的两个子节点
		（3）链接跨越父节点的两个子节点
	二叉树的问题难点在于，如何把题目的要求细化成每个节点需要做的事情
*/
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var dfs func(left, right *Node)
	dfs = func(left, right *Node) {
		if left == nil || right == nil {
			return
		}
		//（1）链接当前左右节点
		left.Next = right
		//（2）链接两个父节点的两个子节点
		dfs(left.Left, left.Right)
		dfs(right.Left, right.Right)
		// (3) 链接跨越父节点的两个子节点
		dfs(left.Right, right.Left)
	}
	dfs(root.Left, root.Right)
	return root
}