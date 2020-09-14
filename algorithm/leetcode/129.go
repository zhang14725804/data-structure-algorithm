/*
	求根到叶子节点数字之和
*/

// 思路1：深度优先遍历获取每个路径，存入数组，然后字符串转数字相加
var res []string

func sumNumbers(root *TreeNode) int {
	dfs(root, strconv.Itoa(root.Val))
	fmt.Println(res)
	sum := 0
	for _, s := range res {
		n, _ := strconv.Atoi(s)
		sum += n
	}
	return sum
}

func dfs(root *TreeNode, s string) {
	if root.Left == nil && root.Right == nil {
		res = append(res, s)
	} else {
		if root.Left != nil {
			dfs(root.Left, s+strconv.Itoa(root.Left.Val))
		}
		if root.Right != nil {
			dfs(root.Right, s+strconv.Itoa(root.Right.Val))
		}
	}
}

// 思路2：每次递归，如果有叶子节点，当前节点值*10,叶子节点位置不变
var ans int
func sumNumbers(root *TreeNode) int {
    if root!=nil{
        dfs(root,0)
    }
    return ans
}
func dfs(root *TreeNode,n int){
	n = n*10 + root.Val
	// 叶子节点
    if root.Left ==nil && root.Right==nil{
        ans += n
    }
    if root.Left!=nil{
        dfs(root.Left,n)
    }
    if root.Right!=nil{
        dfs(root.Right,n)
    }
}