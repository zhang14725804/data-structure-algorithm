package linkedlist

import (
	"fmt"
	"sync"
)

// 节点
type node struct {
	value          int
	previous, next *node
}

// 双向链表
type doublylinkedlist struct {
	mutex *sync.RWMutex // 读写锁
	head  *node
	tail  *node
	size  int
}

func (dl *doublylinkedlist) init() {
	dl.mutex = new(sync.RWMutex)
	dl.size = 0
	dl.head = nil
	dl.tail = nil
}

// 尾插入
func (dl *doublylinkedlist) append(value int) bool {
	no := &node{
		value: value,
	}

	dl.mutex.Lock()
	defer dl.mutex.Unlock()

	if dl.size == 0 {
		dl.head = no
		dl.tail = no
		no.next = nil
		no.previous = nil
	} else {
		no.previous = dl.tail
		no.next = nil
		dl.tail.next = no
		dl.tail = no
	}
	dl.size++
	return true
}

// 查询
func (dl *doublylinkedlist) get(index int) *node {
	if dl.size == 0 || index >= dl.size {
		return nil
	}
	if index == 0 {
		return dl.head
	}
	currentNode := dl.head
	for i := 1; i <= index; i++ {
		currentNode = currentNode.next
	}
	return currentNode
}

// 插入
func (dl *doublylinkedlist) insert(index, value int) bool {
	if index > dl.size {
		return false
	}
	no := &node{
		value: value,
	}
	if index == dl.size {
		return dl.append(value)
	}

	dl.mutex.Lock()
	defer dl.mutex.Unlock()

	if index == 0 {
		no.next = dl.head
		dl.head.previous = no
		dl.head = no
		dl.head.previous = nil
		dl.size++
		return true
	}
	nextNode := dl.get(index)
	no.previous = nextNode.previous
	no.next = nextNode
	nextNode.previous.next = no
	nextNode.previous = no
	dl.size++
	return true
}

// 删除
func (dl *doublylinkedlist) delete(index int) bool {
	if index >= dl.size {
		return false
	}

	dl.mutex.Lock()
	defer dl.mutex.Unlock()
	// 头部
	if index == 0 {
		if dl.size == 1 {
			dl.head = nil
			dl.tail = nil
		} else {
			dl.head.next.previous = nil
			dl.head = dl.head.next
		}
		dl.size--
		return true
	}
	// 尾部
	if index == dl.size-1 {
		dl.tail.previous.next = nil
		dl.tail = dl.tail.previous
		dl.size--
		return true
	}

	no := dl.get(index)
	no.previous.next = no.next
	no.next.previous = no.previous
	dl.size--
	return true
}

// 展示列表
func (dl *doublylinkedlist) display() {
	if dl == nil || dl.size == 0 {
		fmt.Println("this doublylinkedlist is nil or empty")
	}
	dl.mutex.Lock()
	defer dl.mutex.Unlock()

	fmt.Println("this doublylinkedlist`s size is:", dl.size)
	currentNode := dl.head
	for currentNode != nil {
		fmt.Printf("data is %v\n", currentNode.value)
		currentNode = currentNode.next
	}
}

func (dl *doublylinkedlist) reverse() {
	if dl == nil || dl.size == 0 {
		fmt.Println("this doublylinkedlist is nil or empty")
	}
	dl.mutex.Lock()
	defer dl.mutex.Unlock()

	fmt.Println("this doublylinkedlist`s size is:", dl.size)
	currentNode := dl.tail
	for currentNode != nil {
		fmt.Printf("data is %v\n", currentNode.value)
		currentNode = currentNode.previous
	}
}

// DoublyLinkedListTest 测试
func DoublyLinkedListTest() {
	doublylinkedlist := &doublylinkedlist{
		size: 0,
	}
	doublylinkedlist.init()
	doublylinkedlist.append(1)
	doublylinkedlist.append(2)
	doublylinkedlist.append(3)
	doublylinkedlist.append(4)

	doublylinkedlist.display()
	doublylinkedlist.reverse()

	doublylinkedlist.insert(2, 6)
	doublylinkedlist.delete(3)

	doublylinkedlist.display()
}
