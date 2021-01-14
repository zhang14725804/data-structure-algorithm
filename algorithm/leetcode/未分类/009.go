/*
	判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

	思路1：字符串反转之后在和原字符串比较是否相等
	思路2：借助第007题的思路，反转之后在比较
*/
func isPalindrome(x int) bool {
	if x<0{
		return false
	}
	// 保存x
	y:=x
	res:=0
	for x>0{
		res = res*10 + x%10
		x /= 10
	}
	return res == y
}