/*
	0, 1, …, n-1这n个数字(n>0)排成一个圆圈，从数字0开始每次从这个圆圈里删除第m个数字。
	经典约瑟夫问题（递推问题），相邻两项的关系
*/
func lastRemaining(n int, m int) int {
    if n == 1{
		return 0
	}
	// todo：没懂
	return (lastRemaining(n-1, m) + m) % n
}
