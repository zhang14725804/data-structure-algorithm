/*
	1. 长度不想等，直接false
	2. 两字符串相等，如果存在重复元素，true
	3. 两字符串不相等
		a. 判断有且仅有两个s和goal相同位置的字符不一样，一个first, 一个second
		b. second != -1 first!=-1，满足有且仅有两个不同字符s的first等于goal的second且s的second等于goal的first,互换可得相同字符串
*/

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	smap := make(map[byte]int, 0)
	for _, c := range s {
		smap[c]++
	}

	if s == goal {
		for _, ct := range smap {
			if ct > 1 {
				return true
			}
		}
		return false
	}

	first, second := -1, -1
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}
	return first > -1 && second > -1 && s[first] == goal[second] && s[second] == goal[first]
}