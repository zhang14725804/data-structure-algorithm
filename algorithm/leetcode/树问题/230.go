/*
	二叉搜索树中第K小的元素
*/

/*
	方法1：利用BST性质，中序遍历
*/
var nums = 0
var res int

func kthSmallest(root *TreeNode, k int) int {
	inorderTraversal(root, k)
	return res
}

func inorderTraversal(node *TreeNode, k int) {
	// base case
	if node == nil {
		return
	}
	inorderTraversal(node.Left, k)
	nums++
	if nums == k {
		res = node.Val
		return
	}
	inorderTraversal(node.Right, k)
}

/*
	方法2：分治法（todo）

	先计算左子树的节点个数，记为 n，然后有三种情况：

		n 加 1 等于 k，那就说明当前根节点就是我们要找的。
		n加 1 小于 k，那就说明第 k 小的数一定在右子树中，我们只需要递归的在右子树中寻找第 k - n - 1 小的数即可。
		n 加 1 大于 k，那就说明第 k 小个数一定在左子树中，我们只需要递归的在左子树中寻找第 k 小的数即可。
*/
func kthSmallest(root *TreeNode, k int) int {
	n := count(root.Left)
	res := 0
	if n+1 < k {
		res = kthSmallest(root.Right, k-n-1)
	} else if n+1 > k {
		res = kthSmallest(root.Left, k)
	} else {
		res = root.Val
	}
	return res
}

// 统计当前节点有多少个子节点
func count(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return count(node.Left) + count(node.Right) + 1
}