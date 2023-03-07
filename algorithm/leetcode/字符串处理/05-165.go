/*
	1. 以"."为界限分组
	2. 将每组中的字符串转换为数字，比较数字大小
	3. 遍历所有分组
*/
func compareVersion(v1 string, v2 string) int {
	// 遍历两个字符串
	i, j := 0, 0
	for i < len(v1) || j < len(v2) {
		x, y := i, j
		// 找到连续的数字
		for x < len(v1) && v1[x] != '.' {
			x += 1
		}
		for y < len(v2) && v2[y] != '.' {
			y += 1
		}
		// 字符串转数字
		var a int
		if i == x {
			a = 0
		} else {
			a = Str2Int(v1[i:x])
		}
		var b int
		if j == y {
			b = 0
		} else {
			b = Str2Int(v2[j:y])
		}

		// 比较
		if a > b {
			return 1
		}
		if a < b {
			return -1
		}

		// 下一阶段
		i = x + 1
		j = y + 1
	}
	return 0
}

func Str2Int(str string) int {
	res := 0
	for i := 0; i < len(str); i++ {
		res = res*10 + int(str[i]-'0')
	}
	return res
}