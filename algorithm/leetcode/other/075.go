/*
	单指针法（😅😅😅😅）
	两次循环
	1. 第一次把所有的0放在开头
	2. 第二次把所有的1放在0之后
*/
func sortColors(nums []int) {
	cnt0 := swapColors(nums, 0)
	swapColors(nums[cnt0:], 1)
}

func swapColors(colors []int, target int) (countTarget int) {
	for i, c := range colors {
		if c == target {
			colors[i], colors[countTarget] = colors[countTarget], colors[i]
			countTarget++
		}
	}
	return countTarget
}

/*
	双指针法1 😅😅😅
	1. p0, p1分别交换0和1
	2. 遇到0时，
	3. 遇到1时
	TODO 不懂
*/
func sortColors(nums []int) {
	p0, p1 := 0, 0
	for i, c := range nums {
		if c == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			// 😅 向后挪
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			// 😅 p0和p1均向后移动一个位置
			p0++
			p1++
		} else if c == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}