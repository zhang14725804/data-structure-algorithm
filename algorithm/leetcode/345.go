/*
	编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
*/
func reverseVowels(str string) string {
	// careful：先转数组才行，不能直接交换string元素
	s := []byte(str)
	hash := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	// 双指针
	for start, end := 0, len(s)-1; start <= end; {
		if hash[s[start]] && hash[s[end]] {
			s[start], s[end] = s[end], s[start]
			start++
			end--
		} else if !hash[s[start]] {
			start++
		} else if !hash[s[end]] {
			end--
		}
	}
	return string(s)
}