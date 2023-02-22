/*
	我们正在玩一个猜数字游戏。 游戏规则如下：
	我从 1 到 n 选择一个数字。 你需要猜我选择了哪个数字。
	每次你猜错了，我会告诉你这个数字是大了还是小了。
	你调用一个预先定义好的接口 guess(int num)，它会返回 3 个可能的结果（-1，1 或 0）：

	-1 : 我的数字比较小
	1 : 我的数字比较大
	0 : 恭喜！你猜对了！

	思路：二分法
	todo：哪里写错了
*/

func guessNumber(n int) int {
	l, r := 1, n
	for l < r {
		// 这种通常题目描述为满足某种情况的最小的元素。是否存在一个目标值（精确查找）
		mid := (r + l) >> 1
		res := guess(mid)
		if res == 1 {
			r = mid
		} else if res == -1 {
			l = mid + 1
		} else if res == 0 {
			return mid
		}
	}
	return r
}