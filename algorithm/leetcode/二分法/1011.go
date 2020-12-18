
/*
	在 D 天内送达包裹的能力
*/
func shipWithinDays(weights []int, D int) int {
	left, right := getMax(weights), getSum(weights)+1
	// 精确查找，模板一
	for left < right {
		mid := (left + right) >> 1
		if shippable(weights, D, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}