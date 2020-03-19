package main

/*
	栈，先进先出(头插入头删除，或者尾插入尾删除)
	链式栈
*/
type Node struct{
	data interface{}
	Next *Node
}

type LinkedStack interface{
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

func NewStack() *Node {
	//返回一个节点的指针
	return &Node{}
}

func (node *Node) IsEmpty() bool {
	if node.Next == nil {
		return true
	}
	return false
}

func (node *Node) Length() int{
	n := node
	l := 0
	// 循环链表
	for n.Next != nil {
		n = n.Next
		l++
	}
	return l
}

func (node *Node) Push(data interface{}){
	n := &Node{data,nil}
	n.Next = node.Next
	node.Next = n
}

func (node *Node) Pop() interface{}{
	if node.IsEmpty(){
		return nil
	}
	// 要弹出的数据
	val := node.Next.data
	node.Next = node.Next.Next
	return val
}