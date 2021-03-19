/*
	给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
	两棵树重复是指它们具有相同的结构以及相同的结点值。

	ps:前序遍历，中序遍历（唯一对应一棵二叉树，必须加上空节点）
*/

/*
	1、以我为根的这棵二叉树（子树）长啥样？
	2、以其他节点为根的子树都长啥样？

	方法1：
		（1）后序遍历，得到每个节点对应的二叉树（最重要的一步）
		（2）将每个节点对应的二叉树转化为string，存入map
		（3）统计结果
*/
var memory = make(map[string]int)
var ans []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	// 后序遍历
	traverse(root)
	return ans
}

func traverse(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	// 递归遍历左右子树
	left := traverse(root.Left)
	right := traverse(root.Right)
	// 将当前树序列化为字符串
	sub := left + "," + right + "," + strconv.Itoa(root.Val)
	// 存入缓存
	count, _ := memory[sub]
	if count == 1 {
		ans = append(ans, root)
	}
	memory[sub]++
	return sub
}