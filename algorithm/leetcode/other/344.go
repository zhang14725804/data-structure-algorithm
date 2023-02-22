
/*
	方法一：双指针
	start向右走，end向左走
	依次调换位置，直到二者重合
*/
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
	方法二：递归
	TODO 没理解递归
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