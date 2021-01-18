/*
	给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。
	当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
	请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

	说明:
	不允许旋转信封。

	很多算法问题都需要排序技巧，其难点不在于排序本身，而是需要巧妙地排序进行预处理，将算法问题进行转换，为之后的操作打下基础。
	信封嵌套问题就需要先按特定的规则排序，之后就转换为一个 最长递增子序列问题，可以用前文 二分查找详解 的技巧来解决了。

	最长递增子序列问题（Longes Increasing Subsequence，简写为 LIS）

	本体思路：
		先对宽度 w 进行升序排序，如果遇到 w 相同的情况，则按照高度 h 降序排序。之后把所有的 h 作为一个数组，在这个数组上计算 LIS 的长度就是答案。
		这个解法的关键在于，对于宽度 w 相同的数对，要对其高度 h 进行降序排序。
		因为两个宽度相同的信封不能相互包含的，逆序排序保证在 w 相同的数对中最多只选取一个

	todo：这种问题还可以拓展到三维，比如说现在不是让你嵌套信封，而是嵌套箱子，每个箱子有长宽高三个维度，请你算算最多能嵌套几个箱子
*/
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	// 按照宽度升序排列，如果宽度一样，则按照高度降序排列
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if envelopes[j][0] > envelopes[j+1][0] {
				envelopes[j], envelopes[j+1] = envelopes[j+1], envelopes[j]
			} else if envelopes[j][0] == envelopes[j+1][0] {
				// 对于宽度 w 相同的数对，要对其高度 h 进行降序排序
				if envelopes[j][1] < envelopes[j+1][1] {
					envelopes[j], envelopes[j+1] = envelopes[j+1], envelopes[j]
				}
			}
		}
	}

	height := make([]int, n)
	for i := 0; i < n; i++ {
		height[i] = envelopes[i][1]
	}
	return lenOfLIS(height)
}

// 最长递增子序列问题（Longes Increasing Subsequence，简写为 LIS）
func lenOfLIS(nums []int) int {
	piles := 0
	n := len(nums)
	top := make([]int, n)
	for i := 0; i < n; i++ {
		poker := nums[i]
		left, right := 0, piles
		for left < right {
			mid := (left + right) >> 1
			if top[mid] >= poker {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if left == piles {
			piles++
		}
		top[left] = poker
	}
	return piles
}