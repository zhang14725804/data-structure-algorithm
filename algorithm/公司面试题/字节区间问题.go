/*
	给一个排好序的arr(整数列表), 把所有连续的整数合成一个范围，单个数字保持不变, 返回所有的范围集合
     Input: nums = [0,2,3,4,6,8,9]
     Output: ["0","2->4","6","8->9"]

	0110 字节外包岗面试题
*/
func helper(nums []int) []string {
	var res []string
	start, end := 0, 0
	
	for end < len(nums) {
		
		for end < len(nums)-1 && nums[end+1]-1 == nums[end] {
			end++
		}
		
		if start < end {
			res = append(res, sprintf("%d->%d", nums[start], nums[end]))
			end++
			start = end
		} else {
			res = append(res, sprintf("%d", nums[start]))
			start++
			end++
		}
		
	}
	return res
}