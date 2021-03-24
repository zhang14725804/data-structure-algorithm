/*
	二叉树中序序遍历
*/

// 思路1：递归
func inorderTraversal(root *TreeNode) []int {
	// 声明空数组
	res := make([]int, 0)
	// 根节点为空
	if root == nil {
		return res
	}
	// 递归
	dfs(root, &res)
	return res

}
func dfs(node *TreeNode, res *[]int) {
	// 当前节点为空
	if node == nil {
		return
	}
	// 左子树
	dfs(node.Left, res)
	// TODOS：：这个什么语法
	(*res) = append(*res, node.Val)
	// 右子树
	dfs(node.Right, res)
}

// 思路1：递归
var ans []int

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return ans
	}
	inorderTraversal(root.Left)
	ans = append(ans, root.Val)
	inorderTraversal(root.Right)
	return ans
}

/*
	思路2：迭代
	(1)将整棵树的最左边一条链压入栈
	(2)每次取出栈顶元素，如果它有右子树，将右子树压入栈中

	TODO：这样的写法只能遍历根节点，指针和地址的问题😅
	todos：：指针*和地址&
*/

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stk []TreeNode
	p := root
	// 这个他么哪里错了！！！！
	for p != nil || len(stk) > 0 {
		for p != nil {
			stk = append(stk, p)
			p = p.Left
		}
		// top、pop
		p := stk[len(stk)-1]
		res = append(res, p.Val)
		p = p.Right
		stk = stk[:len(stk)-1]
	}
	return res
}

// 迭代方法。这个是对的
func inorderTraversal(root *TreeNode) []int {
	var res []int
	stack := list.New()
	for root != nil || stack.Len() > 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		if stack.Len() > 0 {
			v := stack.Back()
			root = v.Value.(*TreeNode)
			res = append(res, root.Val)
			root = root.Right
			stack.Remove(v)
		}
	}
	return res
}