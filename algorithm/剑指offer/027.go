/*
	数值的整数次方
*/
func Power(base float64, exponent int) float64 {
    res := 1.0
    y := abs(exponent)
	for y > 0 {
		res *= base
		y--
	}
	// 负数的情况
	if exponent < 0{
	    return 1.0/res
	}
	return res
}

func abs(x int) int{
    if x < 0 {
        return -x
    }
    return x
}