/*
	给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以字符串形式返回小数。
	如果小数部分为循环小数，则将循环的部分括在括号内。

	模拟题（难点在于判断循环）。高精度除法问题（高精度加减乘除四类问题）
*/
func fractionToDecimal(numerator int, denominator int) string {
	n,d := numerator,denominator
	// 正负数
	minus:=false
	if n<0{
		minus = !minus
		n = -n
	}
	if d<0{
		minus = !minus
		d = -d
	}
	// number转字符串
	res := strconv.Itoa(n/d)
	n %= d
	// 可以整除
	if n == 0{
		if minus == true && res != "0"{
			return "-"+res
		}
		return res
	}

	// 有余数的情况
	res += "."
	// 记录所有的余数对应的商在小数点后第几位
	hash := make(map[int]int,0)
	// 当前余数
	for n > 0{
		// 如果出现过这个余数
		if hash[n] > 0{
			res = res[0:hash[n]] + "("+res[hash[n]:]+")"
			break
		}
		hash[n] = len(res)
		// 除。。。
		n*=10
		res+= strconv.Itoa(n/d)
		n%=d
	}
	if minus == true{
		res = "-"+res
	}
	return res
}