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
	12.31
*/

// 双向链表节点
type DoubleLinkedListNode struct{
    key int
    value int
    Prev *DoubleLinkedListNode
    Next *DoubleLinkedListNode
}
// 初始化双向链表节点
func initDoubleLinkedListNode(key,value int) *DoubleLinkedListNode{
    return &DoubleLinkedListNode{
        key:key,
        value:value,
    }
}

type LRUCache struct {
    size int // 当前大小
    capacity int // 总容量
    memory map[int]*DoubleLinkedListNode // 缓存
    head *DoubleLinkedListNode
    tail *DoubleLinkedListNode
}


func Constructor(capacity int) LRUCache {
    l:=LRUCache{
        capacity:capacity,
        memory:make(map[int]*DoubleLinkedListNode),
        head:initDoubleLinkedListNode(0,0),
        tail:initDoubleLinkedListNode(0,0),
	}
	// ps 这里需要注意
    l.head.Next = l.tail
    l.tail.Prev = l.head
    return l
}


func (this *LRUCache) Get(key int) int {
    if node,ok:=this.memory[key];ok{
		// 更新头节点
        this.moveToHead(node)
        return node.value
    }else{
        return -1
    }
}


func (this *LRUCache) Put(key int, value int)  {
    if node,ok:=this.memory[key];ok{
        // 存在的情况下覆盖
		node.value = value
		// 更新头节点
        this.moveToHead(node)
    }else{
		node = initDoubleLinkedListNode(key,value)
		// (1) 加缓存
		this.memory[key]=node
		// (2) 更新头节点
		this.addToHead(node)
		// (3) 更新size
		this.size++
		// 超出容量的情况
        if this.capacity<this.size{
			// (1) 更新尾节点
			mnode:=this.removeTail()
			// (2) 删缓存
			delete(this.memory,mnode.key)
			// (3) 更新size
            this.size--
        }
    }
}

// 将节点移动到头 
func (this *LRUCache) moveToHead(node *DoubleLinkedListNode){
	// （1） 删除节点
	this.removeNode(node)
	// （2） 加到头结点
    this.addToHead(node)
}

func (this *LRUCache) addToHead(node *DoubleLinkedListNode){
	// 更新node节点指针
    node.Next = this.head.Next
    node.Prev = this.head
    // question 更新head节点next和以前节点的prev； 顺序不能反
    this.head.Next.Prev = node
    this.head.Next = node
    
}

// 删除一个节点 改变node节点左右两边的指针指向
func (this *LRUCache) removeNode(node *DoubleLinkedListNode){
    node.Prev.Next = node.Next
    node.Next.Prev = node.Prev
}
// 删除尾节点；返回被删除的节点，用来删除缓存
func (this *LRUCache) removeTail() *DoubleLinkedListNode {
    node:=this.tail.Prev
    this.removeNode(node)
    return node
}