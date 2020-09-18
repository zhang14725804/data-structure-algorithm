/*
	一次买卖股票
*/
func maxProfit(prices []int) int {
	res := 0
	// 两次循环直接结算最大值
	for i := 1; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			res = compare(prices[i]-prices[j], res, true)
		}
	}
	return res
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res := 0
	minp := prices[0]
	// 一边扫描一边记录最小值，一遍扫描搞定
	for i := 0; i < len(prices); i++ {
		res = compare(res, prices[i]-minp, true)
		minp = compare(minp, prices[i], false)
	}
	return res
}