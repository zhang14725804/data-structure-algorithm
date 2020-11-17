package main

import "fmt"

//	二叉堆满足堆特性：父节点的键值总是保持固定的序关系于任何一个子节点的键值，且每个节点的左子树和右子树都是一个二叉堆。
//	当父节点的键值总是大于或等于任何一个子节点的键值时为“最大堆”。当父节点的键值总是小于或等于任何一个子节点的键值时为“最小堆
// todo：合并二叉堆
type minHeap struct {
	heapArray []int
	size      int
	maxSize   int
}

// 创建二叉堆
func newMinHeap(maxSize int) *minHeap {
	minHeap := &minHeap{
		heapArray: []int{},
		size:      0,
		maxSize:   maxSize,
	}
	return minHeap
}

// 是否是叶子节点
func (heap *minHeap) isLeaf(index int) bool {
	return index >= (heap.size/2) && index <= heap.size
}

func (heap *minHeap) parentIndex(index int) int {
	return (index - 1) / 2
}

func (heap *minHeap) leftChildIndex(index int) int {
	return 2*index + 1
}
func (heap *minHeap) rightChildIndex(index int) int {
	return 2*index + 2
}

// 插入二叉堆
func (heap *minHeap) insert(item int) error {
	if heap.size >= heap.maxSize {
		return fmt.Errorf("heap is full!")
	}
	// 添加新元素到叶子节点
	heap.heapArray = append(heap.heapArray, item)
	heap.size++
	heap.up(heap.size - 1)
	return nil
}

// 对于最小堆，是删除最小值；todo：删除任意位置的元素（使用最后一个叶子节点元素填充被删除的元素，然后调整）
func (heap *minHeap) remove() int {
	top := heap.heapArray[0]
	// 把堆存储的最后那个节点移到填在根节点处。再从上而下调整父节点与它的子节点
	heap.heapArray[0] = heap.heapArray[heap.size-1]
	heap.heapArray = heap.heapArray[:heap.size-1]
	heap.size--
	heap.down(0)
	return top
}

// 上浮 调整堆
func (heap *minHeap) up(index int) {
	for heap.heapArray[index] < heap.heapArray[heap.parentIndex(index)] {
		heap.swap(index, heap.parentIndex(index))
		index = heap.parentIndex(index)
	}
}

// 构造二叉堆（小顶堆）
func (heap *minHeap) buildMinHeap() {
	// 最后一个非叶子节点开始调整
	for i := ((heap.size) / 2); i >= 0; i-- {
		heap.down(i)
	}
}

// 下沉操作(不好懂😅)
func (heap *minHeap) down(cur int) {
	if heap.isLeaf(cur) {
		return
	}
	smallest := cur
	left := heap.leftChildIndex(cur)
	right := heap.rightChildIndex(cur)
	// 先左边后右边？（question）
	if left < heap.size && heap.heapArray[left] < heap.heapArray[smallest] {
		smallest = left
	}
	if right < heap.size && heap.heapArray[right] < heap.heapArray[smallest] {
		smallest = right
	}
	if smallest != cur {
		heap.swap(cur, smallest)
		heap.down(smallest)
	}
	return
}

func (heap *minHeap) swap(i, j int) {
	heap.heapArray[i], heap.heapArray[j] = heap.heapArray[j], heap.heapArray[i]
}

// func main() {
// 	inputArray := []int{6, 5, 3, 7, 2, 8}
// 	minHeap := newMinHeap(len(inputArray))
// 	for i := 0; i < len(inputArray); i++ {
// 		minHeap.insert(inputArray[i])
// 	}
// 	minHeap.buildMinHeap()
// 	fmt.Println(minHeap.heapArray)
// 	// for i := 0; i < len(inputArray); i++ {
// 	// 	fmt.Println(minHeap.remove())
// 	// }
// }
