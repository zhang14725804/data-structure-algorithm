/*
	一次买卖股票
*/
func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			res = compare(prices[i]-prices[j], res, true)
		}
	}
	return res
}