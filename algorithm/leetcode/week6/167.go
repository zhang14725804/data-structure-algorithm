/*
	167. Two Sum II - Input array is sorted
	双指针算法

	先用暴力算法，是否单调性，优化

	暴力算法，二重循环（注意下标从1开始，而不是0）

	双指针算法不理解

*/
// 暴力算法
func twoSum(numbers []int, target int) []int {
    for i:=0;i<len(numbers);i++{
		for j:=0;j<i;j++{
			if numbers[i]+numbers[j] == target{
				// 返回数组 注意下标从1开始，而不是0
				return []int{j+1,i+1}
			}
		}
	}
	return []int{-1,-1}
}

/*
	双指针算法
	todos：：[0,0,3,4] 0 这组数据测试无法通过
*/ 
func twoSum(numbers []int, target int) []int {
    for j,i:=0,len(numbers)-1;j<len(numbers);j++{
		for i>0 && numbers[i-1]+numbers[j] >=target{
			i--
		}
		if numbers[i]+numbers[j] == target{
			// 返回数组 注意下标从1开始，而不是0
			return []int{j+1,i+1}
		}
	}
	return []int{-1,-1}
}