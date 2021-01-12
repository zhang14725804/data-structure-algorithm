 

/*
	队列，先进先出（头进尾删，或者尾进头删）
	链式队列
*/
type QueueLink struct {
	head *Node
	tail *Node
}
type LinkedQueue interface {
	EnQueue(data interface{})
	DeQueue() interface{}
	Length() int
}

func NewLinkQueue() *QueueLink {
	return &QueueLink{}
}

func (ql *QueueLink) Length() int {
	n := ql.head
	l := 0
	for n.Next != nil {
		n = n.Next
		l++
	}
	return l
}

// 入队（尾部插入）
func (ql *QueueLink) EnQueue(data interface{}) {
	node := &Node{data, nil}
	if ql.head == nil {
		ql.head = node
		ql.tail = node
	} else {
		ql.tail.Next = node
		ql.tail = ql.tail.Next
	}
}

// 出队（头部删除）
func (ql *QueueLink) DeQueue() interface{} {
	if ql.head == nil {
		return nil
	}
	// 记录头部位置
	node := ql.head
	// 只剩一个
	if ql.tail == ql.head {
		ql.tail = nil
		ql.head = nil
	} else {
		ql.head = ql.head.Next
	}
	return node.data
}