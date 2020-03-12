/*
	todos：：堆的操作（如何实现一个堆）：
	（1）查找最大值
	（2）插入一个数
	（3）删除一个树
*/

/*
	692. Top K Frequent Words（找出出现次数最多的K个单词）

	（找出出现次数最多的K个单词）用小根堆来维护出现次数最多的K个单词

	（找出出现次数最少的K个单词）用大根堆来维护出现次数最多的K个单词
	
	（todos：：归并排序，快排要看看）
*/

/*
	实现最小堆参考（https://yangjiahao106.github.io/2019/01/15/golang-%E6%9C%80%E5%A4%A7%E5%A0%86%E5%92%8C%E6%9C%80%E5%B0%8F%E5%A0%86/）
*/
type minHeap struct{
	Element []int
}
// 数组中第一个元素不使用，存放一个小于堆中任何数字的值用于结束循环
func newMinHeap() *minHeap{
	return &minHeap{Element:[]int{math.MinInt64}}
}
// 插入元素就直接将元素增加到堆的末尾，然后进行上浮操作，使二叉堆有序。
func (h *minHeap) push(v int){
	h.Element = append(h.Element, v)
	n := len(h.Element) - 1
	// 上浮
	for ;h.Element[n/2] > v;n /= 2{
		h.Element[n] = h.Element[n/2]
	}

	h.Element[n] = v
}

/*
	删除最小值
	删除最大元素就直接从二叉堆顶端删除，然后进行下沉操作
*/ 
func (h *minHeap) pop() int{
	if len(h.Element) <= 1 {
		return 0
	}
	min := h.Element[1]
	last := h.Element[len(h.Element)-1]

	var i,child int
	for i = 1;i*2<len(h.Element);i = child {
		child = i*2
		if child < len(h.Element) -1 && h.Element[child+1] < h.Element[child]{
			child += 1
		}
		if last > h.Element[child]{
			h.Element[i] = h.Element[child]
		}else{
			break
		}
	}
	h.Element[i] = last
	h.Element = h.Element[:len(h.Element)-1]	
	return min
}

// 获取最小值
func (h *minHeap) mix() int{
	if len(h.Element) > 1{
		return h.Element[1]
	}
	return 0
}