type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
var ans []string
func binaryTreePaths(root *TreeNode) []string {
	var path string
	dfs(root,path)
	return ans
}
func dfs(root *TreeNode,path string)  {
	if root == nil{
		return
	}
	// string可追加，不可修改
	if len(path) > 0{
		path += "->"
	}
	// 数字转字符串，不能直接用string
	path+=strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil{
		ans = append(ans,path)
	}else{
		dfs(root.Left,path)
		dfs(root.Right,path)
	}
}