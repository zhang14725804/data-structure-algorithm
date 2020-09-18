/*
	给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

	[0,0]返回 00，应该是返回0
*/
func largestNumber(nums []int) string {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if fmt.Sprint(nums[j])+""+fmt.Sprint(nums[j+1]) > fmt.Sprint(nums[j+1])+""+fmt.Sprint(nums[j]) {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	res := ""
	for i := len(nums) - 1; i >= 0; i-- {
		res += fmt.Sprint(nums[i])
	}
	return res
}