/*
	输入n个整数，找出其中最小的k个数。
	知识点：大根堆、小根堆
*/
type NodeHeap  []int

func (heap *NodeHeap) push(x int){
	*heap = append(*heap, x)
	// 算是调整堆吧😅
	BubbleSort(*heap)
}
func (h *NodeHeap) pop(){
    old := *h
	n := len(old)
	*h = old[0 : n-1]
	// todo：为什么下面这么写不行
	// 	*h = *h[:len(*h)-1]
}

func getLeastNumbers_Solution(input []int , k int) []int{
	heap := NodeHeap{}
	for _,x := range input{
		heap.push(x)
		if len(heap) > k{
			heap.pop()
		}
	}
	return heap
}
func BubbleSort(nums []int){
	for i := 0; i < len(nums); i++ {
		// 
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1]{
				nums[j],nums[j+1] = nums[j+1],nums[j]
			}
		}
	}
}