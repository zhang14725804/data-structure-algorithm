/*
	多次买卖股票，你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）

	思路(思路妙啊)：每次交易分解成每天都交易，把所有结果为正数的加起来
*/
func maxProfit(prices []int) int {
	res := 0
	for i := 0; i+1 < len(prices); i++ {
		res += compare(0, prices[i+1]-prices[i], true)
	}
	return res
}