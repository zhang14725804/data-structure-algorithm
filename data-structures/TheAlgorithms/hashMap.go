package main

import (
	"fmt"
	"hash/fnv"
)

/*
	https://github.com/TheAlgorithms/Go/blob/master/data-structures/hash-map/hash_map.go
*/
var defaultCapacity uint64 = 1 << 10 // 默认容量，1024

// 节点
type node struct {
	key   interface{}
	value interface{}
	next  *node
}

// hash结构体
type hashMap struct {
	capacity uint64
	size     uint64
	table    []*node
}

// 初始化
func newHashMap() *hashMap {
	return &hashMap{
		capacity: defaultCapacity,
		// node类型的的数组
		table: make([]*node, defaultCapacity),
	}
}

// 根据key获取value
func (hm *hashMap) get(key interface{}) interface{} {
	// 根据key生成hash值，再从table中获取
	node := hm.getNodeByHash(hm.hash(key))

	if node != nil {
		return node.value
	}
	return nil
}

// 存key、value键值对
func (hm *hashMap) put(key interface{}, value interface{}) interface{} {
	return hm.putValue(hm.hash(key), key, value)
}

// 根据key的hash值，存放key、value键值对
func (hm *hashMap) putValue(hash uint64, key interface{}, value interface{}) interface{} {
	// 没有容量
	if hm.capacity == 0 {
		hm.capacity = defaultCapacity
		hm.table = make([]*node, defaultCapacity)
	}
	node := hm.getNodeByHash(hash)
	// 当前node不存在与hash表中
	if node == nil {
		hm.table[hash] = newNode(key, value)
	} else if node.key == key {
		//当前node存在与hash表中
		hm.table[hash] = newNodeWithNext(key, value, node)
		return value
	} else {
		// todo：这他么什么情况
		hm.resize()
		return hm.putValue(hash, key, value)
	}
	hm.size++
	return value
}

// hash表中是否存在当前key
func (hm *hashMap) contains(key interface{}) bool {
	node := hm.getNodeByHash(hm.hash(key))
	return node != nil
}

// 重新分配内存
func (hm *hashMap) resize() {
	// 这什么操作😅
	hm.capacity <<= 1
	temp := hm.table
	hm.table = make([]*node, hm.capacity)

	for i := 0; i < len(temp); i++ {
		node := temp[i]
		if node == nil {
			continue
		}
		hm.table[hm.hash(node.key)] = node
	}
}
func newNode(key interface{}, value interface{}) *node {
	return &node{
		key:   key,
		value: value,
	}
}
func newNodeWithNext(key interface{}, value interface{}, next *node) *node {
	return &node{
		key:   key,
		value: value,
		next:  next,
	}
}

// 根据hash值获取node，用数组下标查
func (hm *hashMap) getNodeByHash(hash uint64) *node {
	return hm.table[hash]
}

// todo：这个不懂😅
func (hm *hashMap) hash(key interface{}) uint64 {
	// New64 返回一个新的64位 FNV-1 hash.Hash 。它的 Sum 方法将以 big-endian 字节顺序排列值
	h := fnv.New64a()
	h.Write([]byte(fmt.Sprintf("%v", key)))
	hashValue := h.Sum64()
	// 这什么骚操作
	return (hm.capacity - 1) & (hashValue ^ (hashValue >> 16))
}

func main() {
	hashMap := newHashMap()

	hashMap.put("test-1", 10)
	fmt.Println(hashMap.get("test-1"))

	hashMap.put("test-1", 20)
	hashMap.put("test-2", 30)
	hashMap.put(1, 40)

	fmt.Println(hashMap.get("test-1"))
	fmt.Println(hashMap.get("test-2"))
	fmt.Println(hashMap.get(1))

	fmt.Println(hashMap.contains(2))
	fmt.Println(hashMap.contains(1))
	fmt.Println(hashMap.contains("test-1"))
}
