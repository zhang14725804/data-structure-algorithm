/*
	给定一个二叉树, 找到该树中两个指定节点的最近公共祖先（Lowest Common Ancestor，简称 LCA）。
*/

/*
	方法1：DFS-递归， 
	0102 此方法没想到 其他方法更懵逼
	后序遍历：根左右
	求最小公共祖先，需要从底向上遍历，那么二叉树，只能通过【后序遍历】（即：回溯）实现从低向上的遍历方式
	(question) 重点理解下面这句话
	前序遍历可以理解为是从上往下，而【后序遍历】是从下往上，就好比从p和q出发往上走，第一次相交的节点就是最近公共祖先
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 😅 base case，终止条件
	if root == nil || p == root || q == root {
		return root
	}
	// 😅 递归左右子树
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 😅 在左子树中没有找到，那一定在右子树中
	if left == nil {
		return right
	}
	// 😅 在右子树中没有找到，那一定在左子树中
	if right == nil {
		return left
	}
	// 😅 不在左子树，也不在右子树，那说明是根节点
	return root
}

/*
	方法2：迭代方式
	(todo)
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {}

/*
	方法3：思路清奇😄（todo）

	root 节点一定是 p 节点和 q 节点的共同祖先，只不过这道题要找的是最近的共同祖先。
	从 root 节点出发有一条唯一的路径到达 p。
	从 root 节点出发也有一条唯一的路径到达 q。
	要找的最近的共同祖先就是第一次出现分叉的时候

	把从 root 到 p 和 root 到 q 的路径找到
	倒着遍历其中一条路径，然后看当前节点在不在另一条路径中，当第一次出现在的时候，这个节点就是我们要找到的最近的公共祖先了
*/
var pPath, qPath, path []*TreeNode
var p, q *TreeNode

func lowestCommonAncestor(root, _p, _q *TreeNode) *TreeNode {
	p, q = _p, _q
	path = append(path, root)
	findPath(root)
	var ppm = make(map[*TreeNode]bool)
	// map 保存其中一条路径
	for _, v := range pPath {
		ppm[v] = true
	}
	// 倒着遍历另一条一条路径，第一次出现交叉的地方就是最近公共祖先
	for i := len(qPath) - 1; i >= 0; i-- {
		// v,k
		if _, ok := ppm[qPath[i]]; ok {
			return qPath[i]
		}
	}
	return root
}

/*
	【递归+回溯】 
	😄😄😄 ，寻找root到某个节点的路径
*/ 
func findPath(root *TreeNode) {
	if root == nil {
		return
	}
	// 找到p节点
	if root == p {
		cp := make([]*TreeNode, len(path))
		copy(cp, path)
		pPath = cp
	}
	// 找到q节点
	if root == q {
		cp := make([]*TreeNode, len(path))
		copy(cp, path)
		qPath = cp
	}

	if root.Left != nil {
		path = append(path, root.Left)
		findPath(root.Left)
		path = path[:len(path)-1]
	}

	if root.Right != nil {
		path = append(path, root.Right)
		findPath(root.Right)
		path = path[:len(path)-1]
	}

}