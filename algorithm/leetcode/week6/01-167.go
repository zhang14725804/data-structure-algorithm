/*
	方法2：双指针算法
	1. start,end相向而行
	2. 求和sum=start+end
	3. sum>target,end向前走
	4. sum<target,start向后走
*/
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			return []int{i + 1, j + 1}
		}
	}
	return []int{-1, -1}
}