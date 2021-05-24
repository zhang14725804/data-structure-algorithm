/*
	二叉树的后序遍历
*/

/*
	方法1：DFS-递归实现
*/
func postorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}
	dfs(root)
	return res
}

/*
	方法2：迭代实现（ question 妙啊 😅😅）
	先序遍历（根左右）-> 根右左 -> 反转数组->左右根
*/ 
func postorderTraversal(root *TreeNode) []int {
    stack :=make([]*TreeNode,0)
    res:=make([]int,0)
	// 判空
    if root ==nil{
        return res
    }
    stack = append(stack,root)
    for len(stack)>0{
        sLen := len(stack)
        for i:=0;i<sLen;i++{
			// 栈，后进先出
            cLen := len(stack)-1
            cur:=stack[cLen]
            stack = stack[:cLen]
            res = append(res,cur.Val)
			// 左子树后入栈
            if cur.Left!=nil{
                stack = append(stack,cur.Left)
            }
			// 右子树先入栈
            if cur.Right!=nil{
                stack = append(stack,cur.Right)
            }
        }
    }
	for i:=0;i<len(res)/2;i++{
        res[i],res[len(res)-i-1] = res[len(res)-i-1],res[i]
    }
    return res
}

/*
	方法2：迭代实现（TODO 😅）
*/
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	if root == nil {
		return nil
	}
	
	// 通过lastVisit标识右子节点是否已经弹出
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 这里先看看，先不弹出
		node := stack[len(stack)-1]
		// 根节点必须在右节点弹出之后，再弹出
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1] // pop
			result = append(result, node.Val)
			// 标记当前这个节点已经弹出过
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}