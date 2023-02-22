/*
	给定一个整数，将其转化为7进制，并以字符串形式输出。

	进制转换，我反应不过来😅(question)
*/
func convertToBase7(num int) string {
	res := ""
	flag := false
	// 正负数处理
	if num < 0 {
		num = -num
		flag = true
	}
	// %取余数，/取商
	for num/7 != 0 {
		res = fmt.Sprintf("%v", num%7) + res // fmt.Sprintf整数转字符串
		num /= 7
	}
	res = fmt.Sprintf("%v", num) + res

	if flag {
		res = "-" + res
	}
	return res
}