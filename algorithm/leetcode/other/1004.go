

/*
	方法一：二分法
	1. 对于任意的右端点 right，希望找到最小的左端点 left，使得 [ left , right ]  包含不超过  k 个 0。
	2. 要想快速判断一个区间内 0 的个数，我们可以考虑将数组 A 中的 0 变成 1， 1 变成 0
	3. 对数组A求出前缀和，记为P

		P[right]−P[left−1]≤k
		P[left−1]≥P[right]−k

	左侧的下标是 left − 1 而不是 left，那么我们在构造前缀和数组时，可以将前缀和整体向右移动一位，空出 P[0] 的位置
		P[0]=0
		P[i]=P[i−1]+(1−A[i−1])
*/
func longestOnes(nums []int, k int) (ans int) {
	nLen := len(nums)
	p := make([]int, n+1)
	// 前缀和，统计0的个数 😅😅😅
	for i, v := range nums {
		p[i+1] = p[i] + 1 - v
	}
	for right, v := range p {
		// 二分查找，找到元素的下标。
		// 😅 如果不存在，返回该插入的位置
		left := binarySearch(p, v-k)
		ans = Max(ans, right-left)
	}
	return
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (right-left)/2 + left
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

/*
	😅😅😅😅
	方法二：滑动窗口
	把【最多可以把 K 个 0 变成 1，求仅包含 1 的最长子数组的长度】转换为 【找出一个最长的子数组，该子数组内最多允许有 K 个 0 】
*/
func longestOnes(nums []int, k int) (ans int) {
	res := 0
	zeros := 0
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			zeros++
		}
		// 收缩窗口
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		res = max(res, right-left+1)
		right++
	}
	return res
}