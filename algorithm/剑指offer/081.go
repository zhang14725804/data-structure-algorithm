/*
	从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的
	
	模拟问题

	（1）删除0
	（2）是否有重复数字
	（3）最大值和最小值差<=4
*/
func isContinuous(nums []int) bool {
    if len(nums) == 0{
		return false
	}
	sorts(nums)
	k := 0
	for nums[k]==0{
	    k++
	}
	for i := k+1; i < len(nums); i++ {
		if nums[i] == nums[i-1]{
			return false
		}
	}
	return nums[len(nums)-1]-nums[k] <= 4
}

// 冒泡
func sorts(nums []int){
	for i := 0; i < len(nums); i++ {
		for k := 0; k < len(nums)-1; k++ {
			if nums[i]<nums[k]{
				nums[i],nums[k] = nums[k],nums[i]
			}
		}
	}
}