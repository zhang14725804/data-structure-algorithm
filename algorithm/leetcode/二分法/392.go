/*
	思路都没有，活该不懂
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

/*
	1. 遍历s，在t中找s[i]
	2. 设置一个idx，标记上一个在t中s[i]的位置
	3. 若在t中找到了s[i]，并且在idx之后，则标记idx，下一轮循环
	4. 若未找到，直接返回

	// s, t := "aaaaaa", "bbaaaa"
	// s, t := "ace", "abcde"
	// s, t := "aec", "abcde"
	// s, t := "axc", "ahbgdc"
	// s="abc",t=""
*/
func isSubsequence(s string, t string) bool {
	// s="abc",t=""
	if len(s) > len(t) {
		return false
	}

	idx := 0
	for i := 0; i < len(s); i++ {
		// t走完，s还未走完的情况
		if idx == len(t)-1 {
			return false
		}
		// 第一个字母之后的
		if idx > 0 {
			idx++
		}
		for j := idx; j < len(t); j++ {
			// 找到，跳出循环
			if t[j] == s[i] {
				idx = j
				break
			}
			// 未找到，返回false
			if j == len(t)-1 {
				return false
			}
		}
	}
	return true
}