/*
	给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
	有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
	例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
*/
var ans []string
var s string

func restoreIpAddresses(_s string) []string {
	var path string
	s = _s
	helper(0, 0, path)
	return ans
}

/*
	u 字符串下标
	k 分组下标（ip由四个分组构成）
*/
func helper(u, k int, path string) {
	// base case 😅😅😅
	if u == len(s) {
		if k == 4 {
			ans = append(ans, path[1:])
			return
		}
	}
	// 不满足条件pass
	if k > 4 {
		return
	}
	// 当前这一位是0，只能是单独的一个数，例如：103.0.169.39
	if s[u] == '0' {
		//
		helper(u+1, k+1, path+".0")
	} else {
		// 23,1；23*10 + 1
		for i, t := u, 0; i < len(s); i++ {
			// 😅 byte转int
			si, _ := strconv.Atoi(string(s[i]))
			t = t*10 + si
			if t < 256 {
				helper(i+1, k+1, path+"."+strconv.Itoa(t))
			} else {
				break
			}
		}
	}
}