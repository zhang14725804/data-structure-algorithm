/*
	todo :不懂,蒙圈
	代码有问题，无非就是golang中byte，rune和string之间的破事
	和093类似
*/ 
func decodeString(s string) string  {
	res:=""
	for i := 0; i < len(s); i++ {
		k:=0
		// 是否为十进制数字字符
		for isdigit(s[i]){
			// 这里有问题
			k = k*10 + s[i++] - '0'
		}
		j,sum := i+1,1
		for sum > 0{
			if s[j] == "["{
				sum++
			}
			if s[j] == "]"{
				sum--
			}
			j++
		}
		// todo
		r:=decodeString(s[i+1:j-i-2])
		for k>0{
			res+=r
			k--
		}
		i=j
	}
	return res
}

// 是否为十进制数字字符
func isdigit(s){
    pattern := "\\d" //反斜杠要转义
    res,_ := regexp.MatchString(pattern,s)

    return res
}