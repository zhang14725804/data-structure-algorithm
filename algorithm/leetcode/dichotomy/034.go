package leetcode

// Leetcode034 Find First and Last Position of Element in Sorted Array（常规方法和二分法）
func Leetcode034() []int {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	n := len(nums)
	// 常规map方法
	// mp := make(map[int][]int)
	// for i := 0; i < n; i++ {
	// 	if arr, ok := mp[nums[i]]; ok {
	// 		arr = append(arr, i)
	// 		mp[nums[i]] = arr
	// 	} else {
	// 		mp[nums[i]] = []int{i}
	// 	}
	// }
	// res, _ := mp[target]
	// // 判空
	// if len(res) == 0 {
	// 	return []int{-1, -1}
	// }
	// l := res[0]
	// r := res[len(res)-1]
	// return []int{l, r}

	// 二分法（两次二分法）
	// 判空
	if n == 0 {
		return []int{-1, -1}
	}
	// 左边
	left, right := 0, n-1
	for left < right {
		// 注意临界判断
		mid := (left + right) >> 1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if nums[right] != target {
		return []int{-1, -1}
	}
	start := right

	// 右边
	left, right = 0, n-1
	for left < right {
		// 注意临界条件
		mid := (left + right + 1) >> 1
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	end := right
	return []int{start, end}
}
