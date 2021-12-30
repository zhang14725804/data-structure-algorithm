/*
	运用你所掌握的数据结构，设计和实现一个  LRU (Least Recently Used 最近最少使用) 缓存机制。
	它应该支持以下操作： 获取数据 get 和 写入数据 put 。
	LRU认为最近使用过的数据应该是是「有用的」，很久都没用过的数据应该是无用的，内存满了就优先删那些很久没用过的数据。

	获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
	写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。
	当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

	进阶:
		你是否可以在 O(1) 时间复杂度内完成这两种操作？
*/

/*
	方法1：哈希表 + 双向链表（todo）
	12.30
*/

// 双向链表节点
type DoubleLinkedListNode struct{
    key int
    value int
    Prev *DoubleLinkedListNode
    Next *DoubleLinkedListNode
}
// ps：在双向链表的实现中，使用一个伪头部（dummy head）和伪尾部（dummy tail）标记界限，
// 这样在添加节点和删除节点的时候就不需要检查相邻的节点是否存在。
type LRUCache struct {
	size       int // 实际大小
	capacity   int // 总容量
	cache      map[int]*DoubleLinkedListNode // 缓存
	head *DoubleLinkedListNode // head表示最近使用
	tail *DoubleLinkedListNode// tail表示最久未使用
}

func initDoubleLinkedListNode(key,value int)*DoubleLinkedListNode{
    return &DoubleLinkedListNode{
        key:key,
        value:value,
    }
}

// 初始化LRU
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		// cache:make(map[int]*DoubleLinkedListNode),
		cache:    map[int]*DoubleLinkedListNode{},
		head:     initDoubleLinkedListNode(0, 0),
		tail:     initDoubleLinkedListNode(0, 0),
		capacity: capacity,
	}
	// 连接头尾
	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	// 判断是否存在
	if node, ok := this.cache[key]; !ok {
		return -1
	}else{
		// 最近使用；添加到头
		this.moveToHead(node)
		return node.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if cnode, ok := this.cache[key]; !ok {
		cnode = initDoubleLinkedListNode(key, value)
		this.cache[key] = cnode
		this.addToHead(cnode)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			// 从hash表中删除
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		// 若key存在，覆盖；添加到头
		cnode.value = value
		this.moveToHead(cnode)
	}
}

// 原来的节点，移动到头部，表示最近使用
func (this *LRUCache) moveToHead(node *DoubleLinkedListNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// 新节点，添加到头部
func (this *LRUCache) addToHead(node *DoubleLinkedListNode) {
	// 改变node节点Prev,Next指针
	node.Prev = this.head
	node.Next = this.head.Next
	// 改变head节点Prev,Next指针
	this.head.Next.Prev = node
	this.head.Next = node
}

// 删除节点
func (this *LRUCache) removeNode(node *DoubleLinkedListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// 删除最后一个元素
func (this *LRUCache) removeTail() *DoubleLinkedListNode {
	node := this.tail.Prev
	this.removeNode(node)
	return node
}