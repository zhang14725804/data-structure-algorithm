/*
	数据流中的中位数（动态求中位数）
	知识点：大根堆，小根堆（todo：如何实现（手写）大小根堆）
	思路：一个大根堆，一个小根堆，中间那个数就是中位数
*/

// 首先定义大小根堆（这里采用数组实现的类大小根堆，意思到了）
type MaxHeap  []int
type MinHeap  []int

var maxHeap  MaxHeap
var minHeap  MinHeap
/*
	每次把元素插入大顶堆，判断maxHeap.size > minHeap.size+1，调整大小堆
*/ 
func insert(num int) {
	maxHeap.maxPush(num)
	if len(minHeap) > 0 && maxHeap.maxTop() > minHeap.minTop(){
		maxv:= maxHeap.maxTop()
		minv:= minHeap.minTop()
		maxHeap.maxPop()
		minHeap.minPop()
		maxHeap.maxPush(minv)
		minHeap.minPush(maxv)
	}
	if len(maxHeap) > len(minHeap) + 1{
		minHeap.minPush(maxHeap.maxTop())
		maxHeap.maxPop()
	}
}

func getMedian() float64{
    // fmt.Println("maxHeap:",maxHeap,"minHeap:",minHeap)
    // fmt.Println("是否是奇数：",(len(maxHeap) + len(minHeap)) & 1 == 1)
    // 判断奇偶数
    if (len(maxHeap) + len(minHeap)) & 1 == 1{
		return float64(maxHeap.maxTop())
	}
// 	fmt.Println(maxHeap.maxTop(),minHeap.minTop())
	return float64(maxHeap.maxTop() + minHeap.minTop())/2.0
}

/*
	大根堆push、top、pop方法
	todo：这么写太笨了😅
*/ 
func (heap *MaxHeap) maxPush(x int){
	*heap = append(*heap, x)
	// 算是调整堆吧😅
	BubbleSort(*heap, true)
}
func (h *MaxHeap) maxPop(){
    old := *h
	n := len(old)
	*h = old[0 : n-1]
}

func (h *MaxHeap) maxTop() int{
	old := *h
	n := len(old)
    return old[n-1]
}

/*
	小根堆push、top、pop方法
	todo：这么写太笨了😅
*/ 
func (heap *MinHeap) minPush(x int){
	*heap = append(*heap, x)
	// 算是调整堆吧😅
	BubbleSort(*heap, false)
}
func (h *MinHeap) minPop(){
    old := *h
	n := len(old)
	*h = old[0 : n-1]
}

func (h *MinHeap) minTop() int{
	old := *h
	n := len(old)
    return old[n-1]
}
// flag true大顶堆，false小顶堆
func BubbleSort(nums []int, flag bool){
	for i := 0; i < len(nums); i++ {
		// 
		for j := 0; j < len(nums)-i-1; j++ {
			if flag{
				if nums[j] > nums[j+1]{
					nums[j],nums[j+1] = nums[j+1],nums[j]
				}
			}else{
				if nums[j] < nums[j+1]{
					nums[j],nums[j+1] = nums[j+1],nums[j]
				}
			}
			
		}
	}
}