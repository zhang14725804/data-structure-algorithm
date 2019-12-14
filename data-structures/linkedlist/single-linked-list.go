package linkedlist

/*
	Append(t) adds an Item t to the end of the linked list
	Insert(i, t) adds an Item t at position i
	RemoveAt(i) removes a node at position i
	IndexOf() returns the position of the Item t
	IsEmpty() returns true if the list is empty
	Size() returns the linked list size
	String() returns a string representation of the list
	Head() returns the first node, so we can iterate on it
*/
import "fmt"

// Node 节点
type Node struct {
	Name string
	Next *Node
}

// NodeList 链表
type NodeList struct {
	Name string
	Head *Node
}

// CreateNodeList 创建
func CreateNodeList(name string) *NodeList {
	return &NodeList{
		Name: name,
	}
}

// AddNode 尾插入
func (p *NodeList) AddNode(name string) error {
	s := &Node{
		Name: name,
	}
	if p.Head == nil {
		p.Head = s
	} else {
		currendNode := p.Head
		for currendNode.Next != nil {
			currendNode = currendNode.Next
		}
		currendNode.Next = s
	}
	return nil
}

// RemoveNode 删除节点
func (p *NodeList) RemoveNode(name string) error {
	currentNode := p.Head
	if currentNode.Name == name {
		h := currentNode.Next
		p.Head = h
	} else {
		for currentNode.Next != nil {
			node := currentNode
			currentNode = currentNode.Next
			if currentNode.Name == name {
				node.Next = currentNode.Next
			}
		}
	}
	return nil
}

// ShowAllNodes 显示所有节点
func (p *NodeList) ShowAllNodes() error {
	currentNode := p.Head
	if currentNode == nil {
		fmt.Println("NodeList is empty")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.Next != nil {
		currentNode = currentNode.Next
		fmt.Printf("%+v\n", *currentNode)
	}
	return nil
}

// SingleLinkedListDemo 测试简单链表
func SingleLinkedListDemo() {
	con := "我的最爱"
	nodeList := CreateNodeList(con)
	nodeList.AddNode("Ophelia")
	nodeList.AddNode("Shape of you")
	nodeList.ShowAllNodes()
	nodeList.RemoveNode("Shape of you")
	nodeList.ShowAllNodes()
}
