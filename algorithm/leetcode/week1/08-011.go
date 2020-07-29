/*
	盛最多水的容器

	思路：双指针法；开始时一个指向头，一个指向尾，那个指向的位置矮（如果高度相同，随机移动），那个向中间移动，移动的过程中记录最大值，直到两个指针相遇位置
	证明过程比较难懂😅
*/
func maxArea(height []int) int {
	res := 0
	for l, r := 0, len(height)-1; l < r; {
		left, right := height[l], height[r]
		res = compare(res, (r-l)*compare(left, right, false), true)
		if left < right {
			l++
		} else {
			r--
		}
	}
	return res
}