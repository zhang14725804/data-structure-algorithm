/*
	要从下往上看，局部最优：让叶子节点的父节点安摄像头，所用摄像头最少，整体最优：全部摄像头数量所用最少！

	确定遍历顺序：后序遍历
	如何隔两个节点放一个摄像头？用三种状态表示每个节点的状态
		0：该节点无覆盖
		1：本节点有摄像头
		2：本节点有覆盖
	因为在遍历树的过程中，就会遇到空节点，那么问题来了，空节点究竟是哪一种状态呢？ 空节点表示无覆盖？ 表示有摄像头？还是有覆盖呢？
		空节点的状态只能是有覆盖，这样就可以在叶子节点的父节点放摄像头了

	单层逻辑处理
		情况1：左右节点都有覆盖
		情况2：左右节点至少有一个无覆盖的情况
		情况3：左右节点至少有一个有摄像头
		情况4：头结点没有覆盖
*/
var res int

func minCameraCover(root *TreeNode) int {
	res = 0
	// 情况4
	if dfs(root) == 0 {
		res++
	}
	return res
}

func dfs(root *TreeNode) int {
	// 空节点，该节点有覆盖
	if root == nil {
		return 2
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	// 情况1
	if left == 2 && right == 2 {
		return 0
	}
	// 情况2
	if left == 0 || right == 0 {
		res++
		return 1
	}
	// 情况3
	if left == 1 || right == 1 {
		return 2
	}
	return -1
}