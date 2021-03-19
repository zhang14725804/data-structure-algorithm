/*
	二叉树的序列化与反序列化
	考察二叉树遍历方式：前序遍历、中序遍历、后序遍历、层序遍历（迭代）
	不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。
	(question):二叉树遍历集大成者🔥🔥🔥

	PS：单单前序遍历结果是不能还原二叉树结构的，因为缺少空指针的信息，至少要得到前、中、后序遍历中的两种才能还原二叉树
*/

/**********************前序遍历做法************************/

// 类
type Codec struct {
	codec []string
}

// 初始化Codec类
func Constructor() Codec {
	return Codec{}
}

// 类方法 Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 前序遍历
	return dfs(root, "")
}
func dfs(root *TreeNode, str string) string {
	if root == nil {
		str += "null,"
	} else {
		str += strconv.Itoa(root.Val) + ","
		str = dfs(root.Left, str)
		str = dfs(root.Right, str)
	}
	return str
}

// 类方法 Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// 根据","分割字符串
	l := strings.Split(data, ",")
	this.codec = l
	// 前序遍历
	return this.rdfs()
}

func (this *Codec) rdfs() *TreeNode {
	// 从前向后
	if this.codec[0] == "null" {
		this.codec = this.codec[1:]
		return nil
	}
	// 左右根的顺序
	val, _ := strconv.Atoi(this.codec[0])
	root := &TreeNode{Val: val}
	this.codec = this.codec[1:]
	root.Left = this.rdfs()
	root.Right = this.rdfs()
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */