package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	105. Construct Binary Tree from Preorder(前序遍历) and Inorder(中序遍历) Traversal
*/

/*
	var pos map[int]int
	这样写会报错：assignment to entry in nil map
*/
var pos = make(map[int]int)

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	for i := 0; i < n; i++ {
		pos[inorder[i]] = i
	}
	return dfs105(preorder, inorder, 0, n-1, 0, n-1)
}

// 递归函数
func dfs105(preorder, inorder []int, pl, pr, il, ir int) *TreeNode {
	if pl > pr {
		return nil
	}
	val := preorder[pl]
	k := pos[val]
	length := k - il
	root := &TreeNode{val, nil, nil}
	root.Left = dfs105(preorder, inorder, pl+1, pl+length, il, k-1)
	root.Right = dfs105(preorder, inorder, pl+length+1, pr, k+1, ir)
	return root
}
