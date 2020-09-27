/*
	数据流的中位数

	方法1：大根堆小根堆，没懂（todo）
	方法2：红黑树😅，更蒙
*/
type MedianFinder struct {
	minHeap *heapMin
	maxHeap *heapMin
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: new(heapMin),
		maxHeap: new(heapMin),
	}
}

// 没懂😅
func (this *MedianFinder) AddNum(num int) {
	if this.minHeap.Len() > 0 && num > (*this.minHeap)[0] {
		heap.Push(this.minHeap, num)
	} else {
		heap.Push(this.maxHeap, -num)
	}

	if this.minHeap.Len()-this.maxHeap.Len() == 2 {
		heap.Push(this.maxHeap, -(heap.Pop(this.minHeap)).(int))
	} else if this.maxHeap.Len()-this.minHeap.Len() == 2 {
		heap.Push(this.minHeap, -(heap.Pop(this.maxHeap)).(int))
	}

}

// 没懂😅
func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() > this.maxHeap.Len() {
		return float64((*this.minHeap)[0])
	} else if this.minHeap.Len() < this.maxHeap.Len() {
		return -float64((*this.maxHeap)[0])
	}
	return float64((*this.minHeap)[0]-(*this.maxHeap)[0]) / float64(2)
}

type heapMin []int

func (h heapMin) Len() int {
	return len(h)
}

func (h heapMin) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *heapMin) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *heapMin) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func (h *heapMin) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */