/*
	请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。

	实现 LFUCache 类：
		LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
		int get(int key) - 如果键存在于缓存中，则获取键的值，否则返回 -1。
		void put(int key, int value) - 如果键已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量时，则应该在插入新项之前，使最不经常使用的项无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最久未使用 的键。

		注意「项的使用次数」就是自插入该项以来对其调用 get 和 put 函数的次数之和。使用次数会在对应项被移除后置为 0 。
*/
/*
	(1)调用get(key)方法时，要返回该key对应的val
	(2)只要用get或者put方法访问一次某个key，该key的freq就要加一
	(3)如果在容量满了的时候进行插入，则需要将freq最小的key删除，如果最小的freq对应多个key，则删除其中最旧的那一个 

	todo:第三步最难实现（🔥🔥🔥）
*/ 
type LinkedHashSet strcut{
}

type LFUCache struct {
	// key到val的映射
	keyToVal map[int]int
	// key到freq的映射
	keyToFreq map[int]int
	// freq到key的映射表 (question)
	freqToKeys map[int]
	// 最小频次
	minFreq int
	// 缓存最大容量
	capacity int 
	// 当前缓存大小
	size int
}

func Constructor(capacity int) LFUCache {

}

func (this *LFUCache) Get(key int) int {

}

func (this *LFUCache) Put(key int, value int) {

}