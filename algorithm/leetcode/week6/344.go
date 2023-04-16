
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
	方法二：递归 😅😅😅
	分解子问题
	TODO 没理解递归
*/
func reverseString(s []byte) {
	res := make([]byte, 0)
	//
	var reverse func(i int)
	reverse = func(i int) {
		// base case
		if i == len(s) {
			return
		}
		reverse(i + 1)
		res = append(res, s[i])
	}

	reverse(0)

	for i := 0; i < len(s); i++ {
		s[i] = res[i]
	}
}
