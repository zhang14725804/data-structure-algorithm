/*
	二叉树中序序遍历
*/

/*
	方法1：DFS-递归实现
	指针、地址 😅😅
*/
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
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
	方法1：DFS-递归实现
*/
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	var dfs func(root *TreeNode)
	dfs=func(root *TreeNode){
		if root == nil {
			return
		}
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return res
}

/*
	思路2：迭代  question 😅😅😅😅
	(1) 将整棵树的最左边一条链压入栈
	(2) 每次取出栈顶元素，如果它有右子树，将右子树压入栈中
*/
func inorderTraversal(root *TreeNode) []int {
    res:=make([]int,0)
    stack:=make([]*TreeNode,0)
    if root==nil{
        return res
    }
    cur:=root
	// 迭代条件 😅
    for cur !=nil || len(stack)>0{
        if cur != nil{
			// 将整棵树的最左边一条链压入栈
            stack = append(stack,cur)
            cur = cur.Left
        }else{
			// 从栈里弹出最后一个元素
            cLen:=len(stack)-1
            cur = stack[cLen]
            stack = stack[:cLen]
			// 根、右
            res = append(res,cur.Val)
            cur = cur.Right
        }
    }
    return res
}