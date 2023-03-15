func minRemoveToMakeValid(s string) string {
	// 记录当前括号的位置，遇到右括号直接存入
	// 遇到左括号弹出一个右，如果没有右括号，记录左括号位置，最后遍历完str，
	cpos := []int{}
	// 记录需要删除的括号的位置
	dpos := []int{}
	for idx, c := range s {
		if c == '(' {
			cpos = append(cpos, idx)
		} else if c == ')' {
			cpl := len(cpos)
			if cpl > 0 {
				cpt := cpos[cpl-1]
				if s[cpt] == '(' {
					cpos = cpos[:cpl-1]
				} else {
					dpos = append(dpos, idx)
				}
			} else {
				dpos = append(dpos, idx)
			}
		}
	}

	// 删除多余的括号，dpos和cpos剩下的部分
	dpos = append(dpos, cpos...)
	hash := make(map[int]struct{}, 0)
	for _, v := range dpos {
		hash[v] = struct{}{}
	}
	bs := []rune(s)
	n := 0
	for i := 0; i < len(bs); i++ {
		if _, ok := hash[i]; ok {
			continue
		}
		bs[n] = bs[i]
		n++
	}
	bs = bs[:n]

	return string(bs)
}