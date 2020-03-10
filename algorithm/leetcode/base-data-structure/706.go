/*
	706. Design HashMap

	hash冲突常见的解决方法：：开放寻址法，拉链法

	开放寻址法：：也就是说当我们存储一个 key，value 时，发现 hashkey (key) 的下标已经被别 key 占用，那我们在这个数组中空间中重新找一个没被占用的存储这个冲突的 key

	拉链法：：简单理解为链表，当 key 的 hash 冲突时，我们在冲突位置的元素上形成一个链表，通过指针互连接，当查找时，发现 key 冲突，顺着链表一直往下找，直到链表的尾节点，找不到则返回空

*/
var N int = 20011
type MyHashMap struct {
	first []int
	second []int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
    return MyHashMap{
		table: make([]int, N),
	}
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
    
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */