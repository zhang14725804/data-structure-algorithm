/*
	1. 遍历邮箱
	2. 根据@的位置切分，前半段是name本地名，后半段是domain
	3. 遇到+号停止，遇到.点跳过
	4. 拼接name@domain存入map
*/
func numUniqueEmails(emails []string) int {
	hash := make(map[string]struct{}, 0)
	for _, email := range emails {
		//
		at := find(email, '@')
		pre := email[:at]
		name := ""
		//
		for _, v := range pre {
			if v == '+' {
				break
			} else if v != '.' {
				name += string(v)
			}
		}
		//
		em := name + "@" + email[at+1:]
		hash[em] = struct{}{}
	}
	return len(hash)
}

func find(arr string, char rune) int {
	for i := 0; i < len(arr); i++ {
		// 注意这里
		if rune(arr[i]) == char {
			return i
		}
	}
	return 0
}