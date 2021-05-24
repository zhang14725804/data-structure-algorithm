/*
	二叉树前序遍历（根左右）
*/

/*
	方法1：DFS-递归实现
*/ 
func preorderTraversal(root *TreeNode) []int {
    var res []int
	// 确定递归函数的参数和返回值
    var dfs func(root *TreeNode)
    // 根左右
    dfs = func(root *TreeNode){
		// base case
        if root ==nil{
            return
        }
		// 单层递归的逻辑
        res = append(res,root.Val)
        dfs(root.Left)
        dfs(root.Right)
    } 
    dfs(root)
    return res
}

/*
	方法2：迭代实现
*/ 
func preorderTraversal(root *TreeNode) []int {
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
			// 右子树先入栈
            if cur.Right!=nil{
                stack = append(stack,cur.Right)
            }
			// 左子树后入栈
            if cur.Left!=nil{
                stack = append(stack,cur.Left)
            }
        }
    }
    return res
}

/*
	迭代实现，标记法
*/
func preorderTraversal(root *TreeNode) []int {
    stack :=make([]*TreeNode,0)
    res:=make([]int,0)
    if root!=nil{
        stack = append(stack,root)
    }
    for len(stack)>0{
		// 弹出顶部元素
        cLen:=len(stack)-1
        cnode := stack[cLen]
        stack = stack[:cLen]
        if cnode !=nil{
			// 右
            if cnode.Right!=nil{
                stack = append(stack,cnode.Right)
            }
			// 左
            if cnode.Left!=nil{
                stack = append(stack,cnode.Left)
            }
			// 根
            stack = append(stack,cnode)
			// question 😅😅😅😅😅😅 中节点访问过，但是还没有处理，加入空节点做为标记。
            stack = append(stack,nil)
        }else{
			// 只有遇到空节点的时候，才将下一个节点放进结果集
            cLen = len(stack)-1
            cnode  =  stack[cLen]
            stack = stack[:cLen]
            res = append(res,cnode.Val)
        }
    }
    return res
}   

