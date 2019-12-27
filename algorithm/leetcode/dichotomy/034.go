package leetcode

// Leetcode034 Find First and Last Position of Element in Sorted Array（常规方法和二分法）
func Leetcode034() []int {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	n := len(nums)
	// 常规map方法
	mp := make(map[int][]int)
	for i := 0; i < n; i++ {
		if arr, ok := mp[nums[i]]; ok {
			arr = append(arr, i)
			mp[nums[i]] = arr
		} else {
			mp[nums[i]] = []int{i}
		}
	}
	res, _ := mp[target]
	// 判空
	if len(res) == 0 {
		return []int{-1, -1}
	}
	l := res[0]
	r := res[len(res)-1]
	return []int{l, r}

	// 二分法（两次二分法）

}
