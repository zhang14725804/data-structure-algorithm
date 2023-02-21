/*
	双指针
	1. start向前走，end倒退
	2. 如果start不是字母，向后走
	3. 如果end不是字母，倒退一步
	4. 如果start!=end，返回false
	5. 双指针向中间靠拢
*/

func isPalindrome(s string) bool {
	start, end := 0, len(s)-1
	bs := []byte(s)
	for start < end {
		for start < end && !isLetter(bs[start]) {
			start++
		}
		for start < end && !isLetter(bs[end]) {
			end--
		}
		if !isLetterEqualCaseless(bs[start], bs[end]) {
			return false
		}
		start++
		end--
	}
	return true
}

func isLetter(s byte) bool {
	return (s >= 97 && s <= 122) || (s >= 65 && s <= 90)
}
func isLetterEqualCaseless(c1, c2 byte) bool {
	return (c1 == c2) || (c1-c2 == 32) || (c2-c1 == 32)
}