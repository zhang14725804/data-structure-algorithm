/*
	给定一个字符串，逐个翻转字符串中的每个单词。

	输入: "the sky is blue"
	输出: "blue is sky the"

	说明：

	无空格字符构成一个单词。
	输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
	如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

	（1）先反转每个单词blue --> eulb
	（2）反转整个字符串
	（3）空格
	todo：这种方法没做出来，从判断k开始有问题
*/
func reverseWords(str string) string {
	s:=[]rune(str)
	// 最终字符串长度
	k:=0
	sLen := len(s)
	// 遍历字符串
	for i:=0;i<sLen;i++{
		// 过滤连续空格
		for i<sLen && string(s[i]) == " " {
			i++
		}
		if i==sLen{
			break
		}
		
		j:=i
		// 连续的非空格
		for j<sLen && string(s[j])!=" " {
			j++
		}
		// 反转单词
		reverse(s[i:j])

		// 
		if k>0{
			// 不能写s[k++]
			k++
			s[k] = ' '
		}
		for i < j-1{
			k++
			i++
			s[k] = s[i]
		}
	}
	// 反转字符串
	reverse(s[0:sLen])
	return string(s)
	
}
func reverse(a []rune) []rune {
    for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }
    return a
}


// 这种方法用了不少库函数
func reverseWords(s string) string {
	// 根据空格拆分字符串
	arr := strings.Split(s, " ")
	res := []string{}
	// 过滤多余空格
	for i := 0; i < len(arr); i++ {
		if arr[i] != " " {
			res = append(res, arr[i])
		}
	}
	// 交换位置
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	/*
		fmt.Sprint(res) 将res数组连同[]转化为字符串类型
		strings.Trim(fmt.Sprint(res), "[]") 去除[]
		func Replace(s, old, new string, n int) string 返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
	*/
	return strings.Replace(strings.Trim(fmt.Sprint(res), "[]"), " ", " ", -1)

}