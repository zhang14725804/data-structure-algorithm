/*
	序列化和反序列化二叉树
	当前用：前序遍历（根左右），实现序列化；序列化简单，反序列化有点难理解
	todo
*/

/**
 * Encodes a tree to a single string.
 */
func serialize(root *TreeNode) string {
	var res string
	dfsEncodes(root, res)
	return res
}

func dfsEncodes(root *TreeNode, res string) {
	if root == nil {
		res += "null "
		return
	}
	res += string(root.Val) + " "
	dfsEncodes(root.Left, res)
	dfsEncodes(root.Right, res)
}

/**
* Decodes your encoded data to tree.
todo:反序列化不太懂
*/
func deserialize(data string) *TreeNode {
	// 起始位置
	u := 0
	return dfsDecodes(data, u)
}

func dfsDecodes(data string, u int) *TreeNode {
	// 遍历完
	if u == len(data) {
		return nil
	}
	fmt.Println(u)
	k := u
	for string(data[k]) != " " {
		k++
	}
	//
	if string(data[u]) == "n" {
		u = k + 1
		return nil
	}
	val := 0
	//
	for i := u; i < k; i++ {
		val = val*10 + int(byte(data[i])-'0')
	}
	u = k + 1
	//
	root := &TreeNode{Val: val, Left: nil, Right: nil}
	//
	root.Left = dfsDecodes(data, u)
	root.Right = dfsDecodes(data, u)
	return root
}