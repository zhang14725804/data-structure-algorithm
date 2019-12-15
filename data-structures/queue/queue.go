package queue

import "fmt"

type element interface{}

/*
	它只允许在表的前端（front）进行删除操作，而在表的后端（rear）进行插入操作，和栈一样，队列是一种操作受限制的线性表。
	进行插入操作的端称为队尾，进行删除操作的端称为队头
*/
type queue interface {
	push(e element) // 向队列添加元素
	pop() element   // 移除队列最前面的元素
	clear() bool    // 清空队列
	size() int      // 获取队列元素个数
	isEmpty() bool  // 判断是否为空
}

type sliceEntry struct {
	element []element
}

// 初始化
func newQueue() *sliceEntry {
	return &sliceEntry{}
}

// 添加元素（尾插入）
func (entry *sliceEntry) push(e element) {
	entry.element = append(entry.element, e)
}

// 移除元素（头删除）
func (entry *sliceEntry) pop() element {
	if entry.isEmpty() {
		fmt.Println("queue is empty")
		return nil
	}
	first := entry.element[0]
	entry.element = entry.element[1:]
	return first
}

// 清空队列
func (entry *sliceEntry) clear() bool {
	if entry.isEmpty() {
		fmt.Println("queue is empty")
		return false
	}

	for i := 0; i < entry.size(); i++ {
		entry.element[i] = nil
	}
	entry.element = nil
	return true
}

func (entry *sliceEntry) isEmpty() bool {
	if len(entry.element) == 0 {
		return true
	}
	return false
}

func (entry *sliceEntry) size() int {
	return len(entry.element)
}

// QueueTest 测试
func QueueTest() {
	queue := newQueue()
	for i := 0; i < 50; i++ {
		queue.push(i)
	}
	fmt.Println("size:", queue.size())
	fmt.Println("移除最前面的元素：", queue.pop())
	fmt.Println("size:", queue.size())
	fmt.Println("清空：", queue.clear())
	for i := 0; i < 50; i++ {
		queue.push(i)
	}
	fmt.Println(queue.pop())
	fmt.Println(queue.pop())
	fmt.Println(queue.size())
}
