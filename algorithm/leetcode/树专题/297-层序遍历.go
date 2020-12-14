type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	str := ""
	if root == nil {
		return str
	}
	// 层序遍历
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		// 出队
		node := queue[0]
		queue = queue[1:]
		// 层序遍历，处理根节点节点
		if node == nil {
			str += "null,"
			continue
		}
		str += strconv.Itoa(node.Val) + ","
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	return str
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	codecs := strings.Split(data, ",")

	queue := make([]*TreeNode, 0)
	root := &TreeNode{Val: stringToInt(codecs[0])}
	queue = append(queue, root)
	for i := 1; i < len(codecs); {
		if len(queue) == 0 {
			break
		}
		// 出队，队列中的节点
		node := queue[0]
		queue = queue[1:]
		// 左节点
		left := codecs[i]
		i++
		if left != "null" {
			node.Left = &TreeNode{Val: stringToInt(left)}
			queue = append(queue, node.Left)
		} else {
			node.Left = nil
		}
		// 右节点
		right := codecs[i]
		i++
		if right != "null" {
			node.Right = &TreeNode{Val: stringToInt(right)}
			queue = append(queue, node.Right)
		} else {
			node.Right = nil
		}
	}
	return root
}

func stringToInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}