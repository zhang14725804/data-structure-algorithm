/*
	运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

	获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
	写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。
	当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

	进阶:
		你是否可以在 O(1) 时间复杂度内完成这两种操作？

	方法1：哈希表 + 双向链表（todo）
*/

// 双向链表
type DoubleLinkedNode struct {
	key, value int
	Prev, Next *DoubleLinkedNode
}

// ps：在双向链表的实现中，使用一个伪头部（dummy head）和伪尾部（dummy tail）标记界限，这样在添加节点和删除节点的时候就不需要检查相邻的节点是否存在。
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DoubleLinkedNode
	head, tail *DoubleLinkedNode
}

func initDoubleLinkedNode(key, value int) *DoubleLinkedNode {
	return &DoubleLinkedNode{
		key:   key,
		value: value,
	}
}
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DoubleLinkedNode{},
		head:     initDoubleLinkedNode(0, 0),
		tail:     initDoubleLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	// 判断是否存在
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key] //
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDoubleLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			// 从hash表中删除
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		// 若key存在，覆盖
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

// 移动到头步，表示最近使用
func (this *LRUCache) moveToHead(node *DoubleLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// 删除节点
func (this *LRUCache) removeNode(node *DoubleLinkedNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// 添加到头步，表示最近使用
func (this *LRUCache) addToHead(node *DoubleLinkedNode) {
	// 改变node节点Prev,Next指针
	node.Prev = this.head
	node.Next = this.head.Next
	// 改变head节点Prev,Next指针
	this.head.Next.Prev = node
	this.head.Next = node
}

// 删除最后一个元素
func (this *LRUCache) removeTail() *DoubleLinkedNode {
	node := this.tail.Prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */