/*
	1. 长度不想等，直接false
	2. 两字符串相等，如果存在重复元素，true
	3. 两字符串不相等，放map中看是否一一匹配

	TODO 未通过测试
*/

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	// s, goal := "abcd", "badc"
	smap := make(map[byte]int, 0)
	gmap := make(map[byte]int, 0)
	sb, gb := []byte(s), []byte(goal)
	for _, c := range gb {
		gmap[c]++
	}
	for _, c := range sb {
		smap[c]++
	}

	if s == goal {
		for _, ct := range smap {
			if ct > 1 {
				// return true
				fmt.Println("true")
			}
		}
		// return false
		fmt.Println("false")
	}

	for c, ct := range smap {
		gct, ok := gmap[c]
		if !ok {
			// return false
			fmt.Println("false")
		}
		if ct != gct {
			// return false
			fmt.Println("false")
		}

	}
	// return true
	fmt.Println("true")
}