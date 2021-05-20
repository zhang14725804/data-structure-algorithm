/*
	编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
	不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
	你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。
*/

/*
	方法1：// 双指针
*/
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
	方法2：递归
*/
func reverseString(s []byte) {
	res := make([]byte, 0)
	reverse(s, 0, &res)
	for i := 0; i < len(s); i++ {
		s[i] = res[i]
	}
}

func reverse(s []byte, i int, res *[]byte) {
	// 递归出口
	if i == len(s) {
		return
	}
	reverse(s, i+1, res)
	*res = append(*res, s[i])
}