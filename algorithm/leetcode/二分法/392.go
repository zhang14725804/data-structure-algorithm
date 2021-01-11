/*
	给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
	字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

	进阶：
	如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

	代码有问题 😅
*/
func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	indexArray := make([][]int, 256)
	for i := 0; i < 256; i++ {
		indexArray[i] = make([]int, 0)
	}
	// 处理t
	for i := 0; i < n; i++ {
		c := t[i]
		indexArray[c] = append(indexArray[c], i)
	}
	// 字符串t上的索引
	j := 0
	for i := 0; i < m; i++ {
		c := s[i]
		// t字符串不包含c字符
		if len(indexArray[c]) == 0 {
			return false
		}
		// 二分搜索区间未包含c
		pos := leftBound(indexArray[c], j)
		if pos == len(indexArray[c]) {
			return false
		}
		// 向前移动指针
		j = indexArray[c][pos] + 1
	}
	return true
}

// 二分查找
func leftBound(arr []int, tar int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := (left + right) >> 1
		if arr[mid] > tar {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}