/*
	求1+2+…+n,要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）
*/
func getSum(n int) int {
	res := n
	// golang语法错误了😅
	(n>0) && (res += getSum(n-1))
	return res
}


// 这么做
func getSum(n int) int {
	var sum int
	// 又是指针问题
    ff(n,&sum)
    return sum
}

func ff(n int, result* int) int{

	*result += n
	// && 相当于if语句
    _ = n>0 && ff(n-1, result)>0
    return n
}