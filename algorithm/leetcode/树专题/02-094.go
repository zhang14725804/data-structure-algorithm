type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	思路1：递归
*/
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

/*
	思路2：迭代
	(1)将整棵树的最左边一条链压入栈
	(2)每次取出栈顶元素，如果它有右子树，将右子树压入栈中

	todo：这样的写法只能遍历根节点，指针和地址的问题😅
	todos：：指针*和地址&
*/

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stk []TreeNode
	p := root
	for p != nil || len(stk) > 0 {
		for p != nil {
			stk = append(stk, p)
			p = p.Left
		}
		// top、pop
		p := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		res = append(res, p.Val)
		p = p.Right
	}
	return res
}

// type Stack []TreeNode

// func (s *Stack) push(node TreeNode) {
// 	*s = append(*s, node)
// }

// func (s *Stack) pop() *TreeNode {
// 	theStack := *s
// 	node := &TreeNode{}
// 	if len(theStack) == 0 {
// 		return node
// 	}
// 	node = &theStack[len(theStack)-1]
// 	*s = theStack[0 : len(theStack)-1]
// 	return node
// }