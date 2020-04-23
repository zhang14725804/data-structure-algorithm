package double_link

import "fmt"

// 双链表结构体
type DoubleLink struct {
	Head *DoubleLinkNode
	Length int
}

func NewDoubleLink() *DoubleLink  {
	return &DoubleLink{
		Head:   NewDoublelinkNode(nil),
		Length: 0,
	}
}

func (list *DoubleLink) Getlength() int  {
	return list.Length
}

func (list *DoubleLink) GetFirstNode() *DoubleLinkNode  {
	return list.Head.Next
}

// 头插入
func (list *DoubleLink) InsertBefore(node *DoubleLinkNode){

	head:=list.Head
	//
	if head.Next == nil{
		node.Next = nil
		head.Next = node
		node.Prev = head
	}else{
		head.Next.Prev = node
		node.Next = head.Next
		head.Next = node
		node.Prev = head
	}
	list.Length++
}

// 尾插入
func (list *DoubleLink) InsertAtfer(node *DoubleLinkNode){
	head:=list.Head
	//
	if head.Next == nil{
		node.Next = nil
		head.Next = node
		node.Prev = head
	}else{
		// 循环找到最后一个
		for head.Next!=nil{
			head = head.Next
		}
		head.Next = node
		node.Prev = head
	}
	list.Length++
}

// 序列化
func (list *DoubleLink) String() string{
	var str1 string
	var str2 string
	head:=list.Head
	for head.Next!=nil{
		str1+=fmt.Sprintf("%v-->",head.Next.Value)
		head = head.Next
	}
	str1+=fmt.Sprintf("nil")
	str1+="\n"
	//
	for head!=list.Head{
		str2+=fmt.Sprintf("<--%v",head.Value)
		head = head.Prev
	}
	str1+=fmt.Sprintf("nil")
	str2+="\n"
	return str1+str2
}

// 节点前插入
func (list *DoubleLink) InsertNodeAfter(dest *DoubleLinkNode, node *DoubleLinkNode) bool {
	p := list.Head
	// 循环遍历链表
	for p.Next != nil && p.Next != dest{
		p = p.Next
	}

	if p.Next == dest{
		if p.Next.Next !=nil{
			p.Next.Next.Prev = node
		}
		node.Next = p.Next.Next
		p.Next.Next = node
		node.Prev = p.Next
		list.Length++
		return true
	}else{
		return false
	}
}

// 节点后插入
func (list *DoubleLink) InsertNodeBefore(dest *DoubleLinkNode, node *DoubleLinkNode) bool{
	p := list.Head
	// 循环遍历链表
	for p.Next != nil && p.Next != dest{
		p = p.Next
	}

	if p.Next == dest{
		if p.Next!=nil{
			p.Next.Prev = node
		}
		node.Next = p.Next
		node.Prev = p
		p.Next = node
		list.Length++
		return true
	}else{
		return false
	}
}
