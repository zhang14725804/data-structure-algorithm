/*
	排序数组
	12.30
	方法1：快速排序（超出时间限制）
	方法2：归并排序
*/

func sortArray(nums []int) []int {
    return quickSort(nums)
}

func quickSort(nums []int)(res []int){
    if len(nums)==0{
        return
	}
	// 选一个基点
    idx:=nums[0]
	left,right:=make([]int,0),make([]int,0)
	// 小的放左边大的放右边
    for i:=1;i<len(nums);i++{
        if nums[i]<idx{
            left=append(left,nums[i])
        }else{
            right=append(right,nums[i])
        }
	}
	// 递归执行左右，再将结果拼起来
    res = append(res,quickSort(left)...)
    res= append(res,idx)
    res = append(res,quickSort(right)...)
    return
}

/*
	12.30 merge写错导致调试40分钟
*/
func sortArray(nums []int) []int {
    return mergeSort(nums)
}
// (1)自上而下的递归
func mergeSort(nums []int) (res []int){
	nLen:=len(nums)
	// base case 递归出口
    if nLen<2{
        return nums
    }
    mid := nLen/2
    left := mergeSort(nums[:mid])
    right := mergeSort(nums[mid:])
    return merge(left,right)
}
// (2) 自下而上的迭代；
func merge(left,right []int) (res []int){
    i,j:=0,0
    for i<len(left) && j<len(right){
        if left[i]<right[j]{
            res=append(res,left[i])
            i++
        }else{
			// 写成 res=append(res,left[j]) 导致结果是相同的n个数字
            res=append(res,right[j])
            j++
        }
    }
    res = append(res,left[i:]...)
    res = append(res,right[j:]...)
    return 
}
