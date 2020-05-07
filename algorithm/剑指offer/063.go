/*
	在字符串中找出第一个只出现一次的字符
	思路，两次for循环，一个hash表
*/
func firstNotRepeatingChar(s string) byte {
	count := make(map[rune]int)
	// cannot use c (type rune) as type byte in map index 😅
	for _,c := range s{
		count[c]++
	}
	res := '#'
	for _,c := range s{
		if count[c] == 1{
			res = c
			break
		}
	}
	return byte(res)
}