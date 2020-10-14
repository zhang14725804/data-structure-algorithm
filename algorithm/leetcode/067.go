/*
	给你两个二进制字符串，返回它们的和（用二进制表示）。
	输入为 非空 字符串且只包含数字 1 和 0。

	卡住了😅
*/
func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	ans := ""
	ca := 0
	for ; i >= 0 || j >= 0; i, j = i-1, j-1 {
		sum := ca
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		ans = fmt.Sprint(sum%2) + ans
		ca = sum / 2
	}
	if ca == 1 {
		ans = fmt.Sprint(ca) + ans
	}
	return ans
}