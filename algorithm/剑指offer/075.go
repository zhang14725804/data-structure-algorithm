/*
	输入一个数组和一个数字s，在数组中查找两个数，使得它们的和正好是s
	一般来说，时间复杂度要比空间复杂度更重要
*/
// 暴力方法
func findNumbersWithSum(nums []int, target int) []int {
    for i:=0;i<len(nums);i++{
		for j:=0;j<i;j++{
			if nums[i] + nums[j] == target{
				return []int{nums[i],nums[j]}
			}
		}
	}
}

// hash表应用
func findNumbersWithSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if hash[target-nums[i]] == 1{
			return []int{target-nums[i],nums[i]}
		}
		hash[nums[i]] = 1
	}
	return []int{}
}