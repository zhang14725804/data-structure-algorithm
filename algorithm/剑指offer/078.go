/*
	左旋转字符串
	字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。

	思路：先反整个字符串，然后分别反转两部分
*/
func leftRotateString(str string, n int) string {
	r := []rune(str)
	reverse(r)
	reverse(r[0:len(str)-n])
	reverse(r[len(str)-n:len(str)])
	return string(r)
}
func reverse(a []rune) {
    for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }
}