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

// 迭代实现前序遍历
var ans []int

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		ans = append(ans, node.Val)
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	return ans
}

