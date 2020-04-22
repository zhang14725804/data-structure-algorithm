package double_link

// 双向链表节点
type DoubleLinkNode struct{
	Value interface{}
	Next *DoubleLinkNode
	Prev *DoubleLinkNode
}

func NewDoublelinkNode(value interface{}) *DoubleLinkNode  {
	return &DoubleLinkNode{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
}

func (node *DoubleLinkNode) GetValue() interface{}  {
	return node.Value
}

func (node *DoubleLinkNode) GetPrev() *DoubleLinkNode  {
	return node.Prev
}

func (node *DoubleLinkNode) GetNext() *DoubleLinkNode  {
	return node.Next
}