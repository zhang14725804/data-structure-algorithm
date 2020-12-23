/*
	给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
	函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

	说明:

	返回的下标值（index1 和 index2）不是从零开始的。
	你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。（数组里面的数，每个只能用一次）
*/
/*
	方法1：暴力算法
*/
func twoSum1(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < i; j++ {
			if numbers[i]+numbers[j] == target {
				// 返回数组 注意下标从1开始，而不是0
				return []int{j + 1, i + 1}
			}
		}
	}
	return []int{-1, -1}
}

/*
	方法2：双指针算法
	i-1 > j. i和j不能相等（不能是同一个数）。 排除[0,0,3,4] 0 情况
*/
func twoSum2(numbers []int, target int) []int {
	// j从左移动，i从右移动，相向而行
	for j, i := 0, len(numbers)-1; j < len(numbers); j++ {
		// 相向而行，具有单调性
		for i-1 > j && numbers[i-1]+numbers[j] >= target {
			i--
		}
		if numbers[i]+numbers[j] == target {
			// 返回数组 注意下标从1开始，而不是0
			return []int{j + 1, i + 1}
		}
	}
	return []int{-1, -1}
}

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