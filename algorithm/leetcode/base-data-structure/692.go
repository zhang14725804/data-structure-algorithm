/*
	方法1: 哈希表 + 排序
		1. 统计单词出现的频率
		2. 单词去重
		3. 相同频率按照字典序输出
*/
func topKFrequent(words []string, k int) []string {
	//
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	uniqueWords := make([]string, 0, len(cnt))
	for w := range cnt {
		uniqueWords = append(uniqueWords, w)
	}
	sort.Slice(uniqueWords, func(i, j int) bool {
		s, t := uniqueWords[i], uniqueWords[j]
		// 首先按照频率，再根据字典序
		return cnt[s] > cnt[t] || cnt[s] == cnt[t] && s < t
	})
	return uniqueWords[:k]
}

/*
	方法2: 优先队列(当前采用此方法)
	如何实现优先队列首先他就是个问题 😅😅😅
*/
func topKFrequent(words []string, k int) []string {
	hash := make(map[string]int)
	heap := newMinHeap()

	for _, word := range words {
		hash[word] += 1
	}

	for key, value := range hash {
		heap.push(&Node{value, key})
		// 注意这里是k+1
		if len(heap.Heap) > k+1 {
			heap.pop()
		}
	}
	// 构造返回结果
	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.pop().second
	}
	return res
}

type minHeap struct {
	Heap []*Node
}

// node(频数，单词)（模拟pair）
type Node struct {
	first  int
	second string
}

/*
	数组中第一个元素不使用，存放一个小于堆中任何数字的值用于结束循环
	todos：：为什么要存一个数，而又不使用呢
*/
func newMinHeap() *minHeap {
	heap := &minHeap{}
	heap.push(&Node{-100000, ""})
	return heap
}

/*
	插入元素就直接将元素增加到堆的末尾，然后进行上浮操作，使二叉堆有序。
	todos：：什么是上浮操作
*/
func (h *minHeap) push(node *Node) {
	h.Heap = append(h.Heap, node)
	n := len(h.Heap) - 1
	// 上浮
	for ; h.Heap[n/2].first > node.first; n /= 2 {
		h.Heap[n] = h.Heap[n/2]
	}
	h.Heap[n] = node
}

/*
	删除最小值
	删除最大元素就直接从二叉堆顶端删除，然后进行下沉操作
	todos：：下沉操作
*/
func (h *minHeap) pop() *Node {
	if len(h.Heap) <= 1 {
		return nil
	}
	last := h.Heap[len(h.Heap)-1]
	// 保存min最小值
	min := h.Heap[1]
	var i, child int
	for i = 1; i*2 < len(h.Heap); i = child {
		child = i * 2
		if child < len(h.Heap)-1 && h.Heap[child+1].first < h.Heap[child].first {
			child += 1
		}
		if last.first > h.Heap[child].first {
			h.Heap[i] = h.Heap[child]
		} else {
			break
		}
	}
	h.Heap[i] = last
	h.Heap = h.Heap[:len(h.Heap)-1]
	return min
}

/*
	方法1: 哈希表 + 排序
*/
func topKFrequent(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	uniqueWords := make([]string, 0, len(cnt))
	for w := range cnt {
		uniqueWords = append(uniqueWords, w)
	}
	sort.Slice(uniqueWords, func(i, j int) bool {
		s, t := uniqueWords[i], uniqueWords[j]
		return cnt[s] > cnt[t] || cnt[s] == cnt[t] && s < t
	})
	return uniqueWords[:k]
}

