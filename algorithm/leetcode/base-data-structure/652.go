/*
	652. Find Duplicate Subtrees

	前序遍历，中序遍历（唯一对应一棵二叉树，必须加上空节点）	

	二叉树序列化之后，再每个字符串存入hashmap，找出出现过两次的字符串，再次hash（有些复杂！）

	todos::做出来困难，做出来再优化更困难
*/

// 每个字符串映射成int
var cnt int = 0
var ans []*TreeNode
var hash = make(map[string]int)
// 经过两部把一棵树变成整数
var count =make(map[int]int)
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    cnt = cnt + 1
	hash["#"] = cnt
	dfs(root)
	return ans
}

func dfs(root *TreeNode) string{
	if root == nil{
		return string(hash["#"])
	}
	left:=dfs(root.Left)
	right:=dfs(root.Right)
	tree := string(root.Val) + ","+left+","+right
	// 
	if count[hash[tree]] == 0  {
        cnt = cnt + 1
		hash[tree] = cnt
	}
	t:=hash[tree]
	count[t]+=1
	if count[t] == 2{
		ans = append(ans,root)
	}
	return string(t)
}