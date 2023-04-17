/*
	方法1:
	交换位置的可能：
	(1)根节点和左子树的某个数字交换 -> 由于根节点大于左子树中的所有数，所以交换后我们只要找左子树中最大的那个数，就是所交换的那个数
	(2)根节点和右子树的某个数字交换 -> 由于根节点小于右子树中的所有数，所以交换后我们只要在右子树中最小的那个数，就是所交换的那个数
	(3)左子树和右子树的两个数字交换 -> 找左子树中最大的数，右子树中最小的数，即对应两个交换的数
	(4)左子树中的两个数字交换
	(5)右子树中的两个数字交换
*/
func recoverTree1(root *TreeNode) {
	if root == nil {
		return
	}
	// 左子树最大节点
	leftMax := getMaxOfTree(root.Left)
	// 右子树最小节点
	rightMin := getMinOfTree(root.Right)
	// 情况3
	if rightMin != nil && leftMax != nil {
		if leftMax.Val > root.Val && rightMin.Val < root.Val {
			temp := rightMin.Val
			rightMin.Val = leftMax.Val
			leftMax.Val = temp
		}
	}
	// 情况1
	if leftMax != nil {
		if leftMax.Val > root.Val {
			temp := root.Val
			root.Val = leftMax.Val
			leftMax.Val = temp
		}
	}
	// 情况2
	if rightMin != nil {
		if rightMin.Val < root.Val {
			temp := root.Val
			root.Val = rightMin.Val
			rightMin.Val = temp
		}
	}
	// 情况4，5
	recoverTree(root.Left)
	recoverTree(root.Right)
}

// 寻找树中最大的节点
func getMaxOfTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := getMaxOfTree(root.Left)
	right := getMaxOfTree(root.Right)
	max := root
	if left != nil && max.Val < left.Val {
		max = left
	}
	if right != nil && max.Val < right.Val {
		max = right
	}
	return max
}

// 寻找树中最小的节点
func getMinOfTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := getMinOfTree(root.Left)
	right := getMinOfTree(root.Right)
	min := root
	if left != nil && min.Val > left.Val {
		min = left
	}
	if right != nil && min.Val > right.Val {
		min = right
	}
	return min
}

/*
	方法2: 中序遍历 （从小到大排列），然后交换位置
	（两种情况：相邻元素（1 3 2 4 5）和不相邻元素（1 5 3 4 2 ））
*/
var first, second, pre *TreeNode

func recoverTree(root *TreeNode) {
	inorderTraversal(root) // 中序遍历
	temp := first.Val
	first.Val = second.Val
	second.Val = temp
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	// 遇到逆序
	if pre != nil && root.Val < pre.Val {
		// 第一次遇到逆序
		if first == nil {
			first = pre
			second = root
		} else {
			// 第二次遇到逆序
			second = root
		}
	}
	//
	pre = root
	inorderTraversal(root.Right)
}