package leetcode
/*
	151. Reverse Words in a String
	方法一
	技巧（todos::技巧懂，但是怎么实现是个问题）
	（1）先反转每个单词blue --> eulb
	（2）反转整个字符串
	（3）空格
*/
func reverseWords(str string) string {
	s:=[]rune(str)
	k:=0
	sLen := len(s)
	for i:=0;i<sLen;i++{
		for i<sLen && string(s[i]) == " " {
			i++
		}
		if i==sLen{
			break
		}
		
		j:=i
		for j<sLen && string(s[j])!=" " {
			j++
		}
		reverse(s[i:j])
		if k>0{
			/*
				很多编程语言都自带前置后置的 ++、-- 运算。但 Go 特立独行，去掉了前置操作，同时 ++、— 只作为运算符而非表达式
				syntax error: unexpected ++, expecting :
				s[k++] = " "
			*/
			k++
			//  cannot use " " (type string) as type rune in assignment
			s[k] = ' '
		}
		for i<j-1{
			k++
			i++
			s[k] = s[i]
		}
	}
	reverse(s[0:sLen])
	return string(s)
	
}
func reverse(a []rune) []rune {
    for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }
    return a
}

/*
	方法二
*/ 
func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	res := []string{}
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			res = append(res, arr[i])
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	// todos::这是赶了个什么
	return strings.Replace(strings.Trim(fmt.Sprint(res), "[]"), " ", " ", -1)

}