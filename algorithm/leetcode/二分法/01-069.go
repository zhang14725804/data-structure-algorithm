/*
	求平方根
	模板2
*/
func mySqrt1(x int) int {
	l,r := 0,x
	for l<r {
		// 
		mid := (l+r+1) >> 1
		// 
		if mid*mid <= x {
			l = mid
		}else{
			r = mid - 1
		}
	}
	return l
}