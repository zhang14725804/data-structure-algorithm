/*
	for大法：
		1. 从前向后遍历order
		2. 从前向后遍历s，交换s中元素的位置，使其和order中的顺序一致
*/
func customSortString(order string, s string) string {
	sb := []rune(s)
	for i := 0; i < len(order); i++ {
		co := order[i]
		start := i
		for j := start; j < len(sb); j++ {
			if sb[j] == rune(co) {
				sb[start], sb[j] = sb[j], sb[start]
				break
			}
		}
	}
	return string(sb)
}

/*
	计数排序（question）
	😅😅😅
*/