/*
	模板1
*/
func mySqrt1(x int) int {
	l,r := 0,x
	for l<r {
		// 
		mid := (l+r+1) >> 1
		if mid*mid <= x {
			l = mid
		}else{
			r = mid - 1
		}
	}
	return l
}

// todo(提交不通过，为什么)
func mySqrt2(x int) int {
	l,r := 0,x
	for l<r {
		// 
		mid := (l+r) >> 1
		if mid*mid > x {
			r = mid - 1
		}else{
			l = mid 
		}
	}
	return l
}