/*
	在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
	排序时间复杂度是 O(NlogN)

	两种方法：
		1. 快速排序之后取倒数第k个即可
		2. 快速选择算法
		3. 二叉堆（优先队列）
*/

// 平平无奇的方法1
func findKthLargest(nums []int, k int) int {
    nums=quickSort(nums)
    return nums[k-1]
}

func quickSort(arr []int) (res []int){
    if len(arr)==0{
        return 
    }
    idx:=arr[0]
    left,right:=make([]int,0),make([]int,0)
    for i:=1;i<len(arr);i++{
        if arr[i]<idx{
            left=append(left,arr[i])
        }else{
            right=append(right,arr[i])
        }
    }
    res=append(res,quickSort(right)...)
    res=append(res,idx)
    res=append(res,quickSort(left)...)
    return 
}

/*
	方法2：部分数据快排
	
	随机选择一个分区点，左边都是大于分区点的数，右边都是小于分区点的数。左部分的个数记做 m。
	如果 k == m + 1，我们把分区点返回即可。
	如果 k > m + 1，说明第 k 大数在右边，我们在右边去寻找第 k - m - 1 大的数即可。
	如果 k < m + 1，说明第 k 大数在左边，我们在左边去寻找第 k 大的数即可。

	输入: [3,2,1,5,6,4] 和 k = 2
	输出: 5
	输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
	输出: 4
*/

var nums []int

func findKthLargest(_nums []int, k int) int {
	nums = _nums
	return helper(0, len(nums)-1, k)
}

func helper(start, end, k int) int {
	// question 采用双指针分区，i 前边始终存比分区点大的元素
	i := start
	// 基准点
	pivot := nums[end] 
	for j := start; j < end; j++ {
		if nums[j] > pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	// ( question ) 将 pivot 值交换到正确的位置
	nums[end], nums[i] = nums[i], nums[end]

	c := i - start + 1
	if c == k {
		// ps： 返回i位置的元素
		return nums[i]
	} else if c < k {
		// 从右边去继续寻找， 😅😅😅 注意参数取值
		return helper(i+1, end, k-c)
	} 
	// 从左边去继续寻找， 😅😅😅 注意参数取值
	return helper(start, i-1, k)
}