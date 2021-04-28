package Single_link

import "fmt"

// 单链表接口
type SingleLink interface {
	GetFirstNode() *SingleLinkNode

	// 头插入尾插入
	InsertFront(node *SingleLinkNode)
	InsertBack(node *SingleLinkNode)
	// 节点头插入尾插入
	InsertBefore(dest interface{},node *SingleLinkNode)
	InsertAfter(dest interface{},node *SingleLinkNode)

	GetNodeAtIndex(index int) *SingleLinkNode
	DeleteNode(node interface{}) bool
	DeleteNodeAtIndex(index int)
	// 获取中间节点（快慢指针 ）
	GetMiddle() *SingleLinkNode
	// 链表反转
	Reverse()
	String() string
}
// 链表struct
type SingleLinkList struct{
	Head *SingleLinkNode
	Length int
}

func NewSingleLinkList() *SingleLinkList {
	return &SingleLinkList{
		Head:   NewSingleLinkNode(nil),
		Length: 0,
	}
}

func (list *SingleLinkList) GetFirst() *SingleLinkNode {
	return list.Head.Next
}
// 头插入
func (list *SingleLinkList) InsertFront(node *SingleLinkNode){
	if list.Head == nil{
		list.Head.Next = node
		node.Next = nil
		list.Length++
	}else{
		//
		back:=list.Head
		node.Next = back.Next
		back.Next = node
		list.Length++
	}
}

// 尾插入
func (list *SingleLinkList) InsertBack(node *SingleLinkNode){
	if list.Head == nil{
		list.Head.Next=node
		node.Next = nil
		list.Length++
	}else{
		//
		p := list.Head
		for p.Next!=nil{
			p = p.Next
		}
		p.Next = node
		list.Length++
	}
}
// 头插入
func (list *SingleLinkList) InsertBefore(dest interface{},node *SingleLinkNode){
	p:=list.Head
	for p.Next!=nil{
		if p.Next.Value == dest{
			node.Next = p.Next
			p.Next = node
			list.Length++
			break
		}
		p = p.Next
	}

}
// 尾插入
func (list *SingleLinkList) InsertAfter(dest interface{},node *SingleLinkNode){
	p:=list.Head
	for p.Next!=nil{
		if p.Next.Value == dest{
			node.Next = p.Next.Next
			p.Next.Next = node
			list.Length++
			break
		}
		p = p.Next
	}
}
// 获取某个位置的元素
func (list *SingleLinkList) GetNodeAtIndex(index int) *SingleLinkNode{
	p:=list.Head
	for index>-1{
		p = p.Next
		index--
	}
	return p
}
// 根据value删除元素
func (list *SingleLinkList) DeleteNode(dest interface{}) bool{
	p := list.Head
	for p.Next!=nil{
		if p.Next.Value == dest{
			p.Next = p.Next.Next
			list.Length--
			return true
		}
		p = p.Next
	}
	return false
}
// 在某个位置删除元素
func (list *SingleLinkList) DeleteNodeAtIndex(index int){
	p:=list.Head
	for index>0{
		p = p.Next
		index--
	}
	p.Next = p.Next.Next
	list.Length--
}

func (list *SingleLinkList) GetMiddle() *SingleLinkNode{
	if list.Head.Next == nil{
		return nil
	}
	p1 := list.Head
	p2 := list.Head
	// 快慢指针
	for p2!=nil && p2.Next!=nil  {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1
}
//
func (list *SingleLinkList) Reverse(){
	if list.Head == nil || list.Head.Next == nil{
		return
	}
	var pre *SingleLinkNode
	current := list.Head.Next
	for current!=nil{
		// 反转
		curNext := current.Next
		current.Next= pre
		pre = current
		current = curNext
	}
	// 
	list.Head.Next = pre
}
// list序列化
func (list *SingleLinkList) String() string{
	var str string
	p:=list.Head
	for p.Next!=nil{
		str+=fmt.Sprintf("%v-->",p.Next.Value)
		p=p.Next
	}
	str+=fmt.Sprintf("nil")
	return str
}