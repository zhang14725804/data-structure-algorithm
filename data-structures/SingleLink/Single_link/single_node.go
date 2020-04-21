package Single_link

// 链表的节点
type SingleLinkNode struct{
	Value interface{}
	Next *SingleLinkNode
}

func NewSingleLinkNode(data interface{}) *SingleLinkNode {
	return &SingleLinkNode{data,nil}
}

func (node *SingleLinkNode) GetValue() interface{} {
	return node.Value
}

func (node *SingleLinkNode) NextNode() *SingleLinkNode {
	return node.Next
}