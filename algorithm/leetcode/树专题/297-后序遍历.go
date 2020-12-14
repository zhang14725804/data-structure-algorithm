/**********************后续遍历做法（注意遍历方向：从后向前，根右左）************************/

type Codec struct {
	codec []string
}

func Constructor() Codec {
	return Codec{}
}

// 类方法 Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 后序遍历
	return dfs(root, "")
}

func dfs(root *TreeNode, str string) string {
	if root == nil {
		str += "null,"
	} else {
		str = dfs(root.Left, str)
		str = dfs(root.Right, str)
		str += strconv.Itoa(root.Val) + ","
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
	// care：从后向前
	if this.codec[len(this.codec)-1] == "null" {
		this.codec = this.codec[:len(this.codec)-1]
		return nil
	}
	// 从后向前在列表取元素（根右左）
	val, _ := strconv.Atoi(this.codec[len(this.codec)-1])
	this.codec = this.codec[:len(this.codec)-1]
	root := &TreeNode{Val: val}
	// 先右后左的顺序
	root.Right = this.rdfs()
	root.Left = this.rdfs()
	return root
}