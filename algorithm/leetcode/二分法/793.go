/*
	阶乘函数后K个零
	K是范围在 [0, 10^9] 的整数。（这就需要认真读题，看看题目给的数据范围有多大。）

	两次二分法(question)（还是不懂二分法 😅）
*/
func preimageSizeFZF(K int) int {
	return rightBound(K) - leftBound(K) + 1
}

func leftBound(target int) int {
	left, right := 0, (1 << 32)
	for left < right {
		// 满足某种情况的最小的元素
		mid := (left + right) >> 1
		if trailingZeroes(mid) < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func rightBound(target int) int {
	left, right := 0, (1 << 32)
	for left < right {
		// 满足某种情况的最大的元素
		mid := (left + right + 1) >> 1
		if trailingZeroes(mid) > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}

// 给定一个整数 n，返回 n! 结果尾数中零的数量
func trailingZeroes(n int) int {
	res := 0
	for n > 0 {
		res += n / 5
		n = n / 5 // 取整
	}
	return res
}