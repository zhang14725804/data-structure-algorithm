/*
	76. Minimum Window Substring

	（1）暴力方法
	（2）滑动窗口（不好理解），双指针问题

	todos::代码有问题
*/
func minWindow(s1 string, t1 string) string {
	// var hash map[rune]int 没有初始化
	hash := make(map[rune]int,0)
	s:=[]rune(s1)
	t:=[]rune(t1)
	for _,c :=range t{
		hash[c]+=1
	}
	// 统计hash表长度（办法好像有点蠢）
	cnt:=0
	for _,_ = range hash{
		cnt+=1
	}

	var res []rune
	// 后一个指针i，走在前面的
	for i,j,c:=0,0,0;i<len(s);i++{
		// 当前字母个数是1
		if hash[s[i]] == 1 {
			c+=1
		}

		hash[s[i]]-=1
		// 当前字母无关紧要
		// 前一个指针j，走在后面的
		for hash[s[j]]<0{
			j+=1
			hash[s[j]]+=1
		}
		// 所有字母都满足要求，更新答案
		if c == cnt{
			if len(res) == 0  || len(res)>i-j+1{
				res=s[j:i-j+1]
			}
		}
	}
	return string(res)
}