/*
	数字序列中某一位的数字
	todo：数位统计问题（和56类似）😅
	todo：测试13通过，10不通过
*/
func digitAtIndex(n int) int {
	// 第一步，确定是几位数
	i:=1 // 
	s:=9 // 最高位
	base:=1 // 起始位
	for n > i*s{
		n -= i*s
		i++
		s *= 10
		base *= 10
	}
	// 第二步，确定是几位数的第几个数子
	number := base + (n+i-1)/i-1
	var r int
	if n%i == 0{
		r = i
	}else{
		r = n%i
	}

	// 第三步，属于那个数的第几个数
	for j:=0;i<i-r;j++{
		number /= 10
	}
	return number%10
}
