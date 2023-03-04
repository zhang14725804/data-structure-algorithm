/*
	1. map统计各个数字出现的次数
	2. uniqNums 统计出现的数字
	3. 根据每个数字出现的次数，对uniqNums进行排序，取前k个即可
*/
func topKFrequent(nums []int, k int) []int {
	cnt := map[int]int{}
	for _, n := range nums {
		cnt[n]++
	}
	uniqNums := make([]int, 0)
	for v := range cnt {
		uniqNums = append(uniqNums, v)
	}
	sort.Slice(uniqNums, func(i, j int) bool {
		s, t := uniqNums[i], uniqNums[j]
		// 按照频率排序序
		return cnt[s] > cnt[t]
	})
	return uniqNums[:k]
}

/*
	方法1：小顶堆,思想好理解，主要是指针地址搞不懂🔥🔥🔥(question)
*/
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




