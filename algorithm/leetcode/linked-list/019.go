package leetcode

// Leetcode019 Remove Nth Node From End of List
func Leetcode019() {
	/*
		单链表，每个节点由data和指向下一个节点的指针组成

		（1）建立虚拟头节点（省去判断是否是头节点）
		（2）让某个指针（假设first）向后走n步
		（3）first，second两个指针同时走，相差n步，当first走到结尾时，终止。second走到了倒数第n+1个节点
		（4）倒数n+1个节点直接指向n的下一个节点就好了
	*/
}
