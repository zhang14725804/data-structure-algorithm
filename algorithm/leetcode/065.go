/*
	验证给定的字符串是否可以解释为十进制数字。

	先设定numSeen，dotSeen和eSeen三种boolean变量，分别代表是否出现数字、点和E
	然后遍历目标字符串
	1.判断是否属于数字的0~9区间
	2.遇到点的时候，判断前面是否有点或者E，都需要return false
	3.遇到E的时候，判断前面数字是否合理，是否有E，并把numSeen置为false，防止E后无数字
	4.遇到-+的时候，判断是否是第一个，如果不是第一个判断是否在E后面，都不满足则return false
	5.其他情况都为false
	最后返回numSeen的结果即可

*/
func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	numSeen, dotSeen, eSeen := false, false, false
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		if arr[i] >= '0' && arr[i] <= '9' {
			numSeen = true
		} else if arr[i] == '.' {
			if dotSeen || eSeen {
				return false
			}
		} else if arr[i] == 'e' || arr[i] == 'E' {
			if eSeen || !numSeen {
				return false
			}
			eSeen = true
			numSeen = false
		} else if arr[i] == '+' || arr[i] == '-' {
			if i != 0 && arr[i-1] != 'e' && arr[i] != 'E' {
				return false
			}
		} else {
			return false
		}
	}
	return numSeen
}