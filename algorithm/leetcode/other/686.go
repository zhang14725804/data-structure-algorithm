/*
	1. acnt,bcnt 统计a、b中出现的字符
	2. 若b中存在a中没有的字符，直接return
	3. 重复a串，使其长度为a+b的1～3倍
	4. 在此期间检查是否有无符合条件的可能
	5. 否则retrun
*/
func repeatedStringMatch(a string, b string) int {
	// 检查b中是否有a中没有的字符
	aset := make(map[rune]struct{}, 0)
	for _, v := range a {
		aset[v] = struct{}{}
	}
	for _, v := range b {
		if _, ok := aset[v]; !ok {
			return -1
		}
	}

	// 备份a、b
	aBack := a
	bBack := b

	cnt := 1 // 默认一次
	// 边界条件 😅不能直接 3*(len(a)+len(b))
	for len(a) < 3*(len(aBack)+len(bBack)) {
		if strings.Contains(a, b) {
			return cnt
		}
		cnt++
		// 不能直接 a+=a
		a += aBack
	}
	return -1
}