package leetcode

import (
	"data-structure-algorithm/algorithm/common"
	"fmt"
)

// Leetcode300 最长上升子序列
func Leetcode300() int {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	n := len(nums)
	f := make([]int, n)

	for i := 0; i < n; i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			// 倒数第二个小于最后一个
			if nums[j] < nums[i] {
				f[i] = common.Max(f[i], f[j]+1)
			}
		}
	}
	var res int
	for k := 0; k < n; k++ {
		res = common.Max(res, f[k])
	}
	fmt.Println(res)
	return res
}
