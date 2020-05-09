/*
	输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变
	思路：先反转整个字符串，再反转单词
*/

func reverseWords(s string) string {
	// 先反转整个字符串
	r := []rune(s)
    r = reverse(r)
	// 再反转单词
	for i:=0; i<len(r); i++{
		j:=i
		for j<len(r) && r[j]!=' '{
			j++
		}
		reverse(r[i:j])
		i = j
	}
	return string(r)
}

func reverse(a []rune) []rune {
    for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }
    return a
}