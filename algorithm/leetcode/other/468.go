/*
	没有技术含量的体力活
	ip := "192.168.1.0"
	ip = "00.0.0.0"
	ip = "256.256.256.256"
	ip = "172.16.254.1"
	ip = "192.168.01.1"
	ip = "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	ip = "2001:db8:85a3:0:0:8A2E:0370:7334"
	ip = "2001:0db8:85a3::8A2E:037j:7334"
	ip = "02001:0db8:85a3:0000:0000:8a2e:0370:7334"
	"2001:0db8:85a3:0:0:8A2E:0370:7334:"
	"1.1.1.1."
	"1.0.1."
	"00.0.0.0"
	"12..33.4"
	"20EE:Fb8:85a3:0:0:8A2E:0370:7334:12"
	"192.0.0.1"
	"222.2f.22.1"
*/
func validIPAddress(queryIP string) string {
	if strings.Contains(queryIP, ".") {
		// 有效的IPv4地址 是 “x1.x2.x3.x4” 形式的IP地址。 其中 0 <= xi <= 255 且 xi 不能包含 前导零
		ip4 := strings.Split(queryIP, ".")
		// ipv4长度要等于4
		if len(ip4) != 4 {
			return "Neither"
		}
		ten := 10
		for _, aip := range ip4 {
			// 不能有空
			if len(aip) == 0 {
				return "Neither"
			}
			// 前导0的情况
			if aip[0] == '0' && len(aip) > 1 {
				return "Neither"
			}
			ipNum := 0
			for _, v := range aip {
				// 每个字符0～9范围内，且数字范围0～255
				if v < '0' || v > '9' {
					return "Neither"
				}
				ipNum = ipNum*ten + int((v - '0'))
			}
			if ipNum > 255 {
				return "Neither"
			}

		}
		return "IPv4"
	}

	// 1 <= xi.length <= 4
	// xi 是一个 十六进制字符串 ，可以包含数字、小写英文字母( 'a' 到 'f' )和大写英文字母( 'A' 到 'F' )。
	// 在 xi 中允许前导零。
	ipv68 := strings.Split(queryIP, ":")
	// 长度为8
	if len(ipv68) != 8 {
		return "Neither"
	}
	for _, ipv6 := range ipv68 {
		// 单个节长度小于4且不为空
		if len(ipv6) > 4 || len(ipv6) == 0 {
			return "Neither"
		}
		// 0～9,a~f,A~F
		for _, w := range ipv6 {
			if (w >= '0' && w <= '9') || (w >= 'a' && w <= 'f') || (w >= 'A' && w <= 'F') {
				continue
			} else {
				return "Neither"
			}
		}
	}
	return "IPv6"
}