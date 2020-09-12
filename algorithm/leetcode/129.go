/*
	求根到叶子节点数字之和

	思路：深度优先遍历获取每个路径，存入数组，然后字符串转数组相加
*/
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