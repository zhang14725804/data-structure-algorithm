/*
	给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
	有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
	例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
*/

/*
	方法1：回溯+判断字符串是合法ip段
	0110 墨迹了一个小时才写完
*/
var ans []string
var path string
var str string

func restoreIpAddresses(s string) []string {
	if len(s) > 12 {
		return ans
	}
	// 只是为了提交，leetcode提交时，ans 会拼接之前提交的结果
	ans = make([]string,0)
    path = ""
    str = s
	backtrack(0, 0)
	return ans
}

/*
   start: 搜索的起始位置
   pointNum: 添加逗点的数量
*/
func backtrack(start int, pointNum int) {
	// 😅😅 base case 逗点数量为3时，分隔结束；判断第四段子字符串是否合法，如果合法就放进path中
	if pointNum == 3 && valid(start, len(str)-1) {
		// 😅😅 这里要拼接最后一步的合法字符串
		path += str[start:len(str)]
		ans = append(ans, path)
		return
	}
	for i := start; i < len(str); i++ {
		// 判断 [start,i] 这个区间的子串是否合法
		if valid(start, i) {
			oLen := len(path)
			// 😅 注意slice取值，左闭右开，所以这里取【i+1】
			path += str[start:i+1] + "."
			pointNum++
			// 😅 这里是i+1,如果写成i会死循环
			backtrack(i+1, pointNum)
			// 😅 回溯
			pointNum--
			path = path[:oLen]
		} else {
			continue
		}
	}
}

/*
	😅 判断字符串s在【左闭右闭】区间[start, end]所组成的数字是否合法
	（1）段位以0为开头的数字不合法
	（2）段位里有非正整数字符不合法
	（3）段位如果大于255了不合法
*/
func valid(start, end int) bool {
	if start > end {
		return false
	}
	// 0开头的数字不合法
	if str[start] == '0' && start != end {
		return false
	}
	num := 0
	// 😅😅😅 包括start、end两个端点，所以这里【i<=end】
	for i := start; i <= end; i++ {
		// 非数字不合法
		if str[i] > '9' || str[i] < '0' {
			return false
		}
		num = num*10 + int(str[i]-'0')
		// 大于255不合法
		if num > 255 {
			return false
		}
	}
	return true
}