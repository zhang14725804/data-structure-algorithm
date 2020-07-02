/*
	A 的旋转操作就是将 A 最左边的字符移动到最右边。 例如, 若 A = 'abcde'，在移动一次之后结果就是'bcdea' 。如果在若干次旋转操作之后，A 能变成B，那么返回True

	思路：将两个a字符串拼接起来组成str，b的所有可能都在str中，只需要判断str中是否存在b
*/

func rotateString(A string, B string) bool {
	if len(A) != len(b){
		return false
	}
	str:=A+A
	// 判断字符串是否包含
	return strings.Contains(str, B)
}