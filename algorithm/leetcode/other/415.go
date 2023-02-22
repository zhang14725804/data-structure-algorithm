func addStrings(num1 string, num2 string) string {
    var res string
    var carry int
	i,j:=len(num1)-1,len(num2)-1
	// 注意循环条件
    for i>=0 || j>=0 || carry!=0{
		var x,y int
		// i，j 字符串转整数
        if i>=0{
            x = int(num1[i]-'0')
            i--
        }
        if j>=0{
            y = int(num2[j]-'0')
            j--
		}
		// 加法
        result := x+y+carry
		res = strconv.Itoa(result%10) + res
		// 进位
        carry = result/10
    }
    return res
}