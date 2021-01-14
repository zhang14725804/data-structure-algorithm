/*
	给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

	你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
	你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
	题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
	你可以按任意顺序返回答案。
*/

// 方法1：小顶堆,思想好理解，主要是指针地址搞不懂🔥🔥🔥(question)
func topKFrequent(nums []int, k int) []int {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num]++
	}

	h := &IHeap{}
	heap.Init(h)
	for key, value := range hash {
		heap.Push(h, &Pair{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(h).(*Pair).Num
	}
	return res
}

// heap 堆的实现 go语言版
type Pair struct {
	Num   int
	Count int
}

type IHeap []*Pair

func (h IHeap) Len() int {
	return len(h)
}

// 这里决定 大小顶堆 现在是小顶堆
func (h IHeap) Less(i, j int) bool {
	return h[i].Count < h[j].Count
}

// 这里为什么不用指针(question)
func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 用指针
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(*Pair))
}

// 用指针
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}




