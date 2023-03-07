/*
	1. k记录字符串最终长度
	2. 遍历字符串
	3. 反转每个单词
	4. 在每个单词后添加空格
	5. 确定每个字符最终位置
	6. 去掉多余的部分
	7. 反转整个字符串

	还是不会 😅😅😅
*/
func reverseWords(str string) string {
	bstr := []rune(str)
	// 最终字符串长度
	k := 0
	sLen := len(str)
	// 遍历字符串
	for i := 0; i < sLen; i++ {
		// 过滤连续空格
		for i < sLen && bstr[i] == ' ' {
			i++
		}
		if i == sLen {
			break
		}

		// 反转单词
		j := i
		for j < sLen && bstr[j] != ' ' {
			j++
		}
		reverse(bstr[i:j])

		// 增加一个空格
		if k > 0 {
			bstr[k] = ' '
			k++
		}

		// 移动单个字符，确定每个单子最终位置
		for i < j {
			bstr[k] = bstr[i]
			k++
			i++
		}
	}

	// 去掉多余的部分
	res := bstr[:k]
	// 反转整个字符串
	reverse(res)

	return string(res)
}

func reverse(a []rune) []rune {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}