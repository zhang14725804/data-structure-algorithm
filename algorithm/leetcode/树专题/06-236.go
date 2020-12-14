/*
	给定一个二叉树, 找到该树中两个指定节点的最近公共祖先（Lowest Common Ancestor，简称 LCA）。
*/

/*
	方法1：递归：本质上是二叉树的【后序遍历】
	(question)前序遍历可以理解为是从上往下，而后序遍历是从下往上，就好比从p和q出发往上走，第一次相交的节点就是最近公共祖先
*/
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	// base case
	if root == nil || p == root || q == root {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//在左子树中没有找到，那一定在右子树中
	if left == nil {
		return right
	}
	//在右子树中没有找到，那一定在左子树中
	if right == nil {
		return left
	}
	// 不在左子树，也不在右子树，那说明是根节点
	return root
}

/*
	方法2：迭代方式(todo)
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
func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	parent = make(map[*TreeNode]*TreeNode, 0) // key：节点，value：节点的父节点
	// 迭代的方式遍历树
	// 将遍历过程中每个节点的父节点保存起来
	// 倒着还原p的路径，并将其放入set中
	// 倒着遍历q的路径，判断是否在p的路径中
}