package leetcode
/*
	151. Reverse Words in a String
	方法一
	技巧（todos::技巧懂，但是怎么实现是个问题）
	（1）先反转每个单词blue --> eulb
	（2）反转整个字符串
	（3）空格
*/
func reverseWords(s string) string {
	k:=0
	sLen := len(s)
	for i:=0;i<sLen;i++{
		for i<sLen && s[i] == " " {
			i++
		}
		if i==sLen{
			break
		}
		
		j:=i
		for j<sLen && s[j]!=" " {
			j++
		}
		reverse(s[i:j])
		if k>0{
			s[k++] = " "
		}
		for i<j{
			s[k++] = s[i++]
		}
	}
	// 去除前后空格
	strings.Trim(k,sLen)
	reverse(s[0,sLen])
	return s
}

// 反转字符串
func reverse(s string) string {
    a := []rune(s)
    for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }
    return string(a)
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