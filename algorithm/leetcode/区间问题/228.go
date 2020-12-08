/*
	给定一个无重复元素的有序整数数组，返回数组区间范围的汇总。
*/

// 方法1：双指针暴力破解
func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return res
	}

	start, end := nums[0], nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1]-1 {
			end++
		} else {
			buildRes(&res, start, end)
			start, end = nums[i+1], nums[i+1]
		}
	}
	buildRes(&res, start, end)
	return res
}

// 操作指针地址😄
func buildRes(res *[]string, start, end int) {
	if start == end {
		*res = append(*res, strconv.Itoa(start))
	} else {
		*res = append(*res, strconv.Itoa(start)+"->"+strconv.Itoa(end))
	}
}

/*
	方法2：借用二分的思想（todo）

	可以一半一半的考虑，比如 1 2 3 4 5 7。先考虑左半部 1 2 3 是否连续，只需要判断下标之差和数字之差是否相等。
	2 - 0 == 3 - 1，所以左半部分是连续的数字，得到一个范围 1 -> 3，而不需要向解法一那样一个一个数字的遍历。
*/ 