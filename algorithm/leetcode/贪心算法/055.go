/*
	😅

	如果某一个作为 起跳点 的格子可以跳跃的距离是 3，那么表示后面 3 个格子都可以作为 起跳点
	可以对每一个能作为 起跳点 的格子都尝试跳一次，把 能跳到最远的距离 不断更新
	如果可以一直跳到最后，就成功了
*/
func canJump(nums []int) bool {
	// 能跳的最远距离
	k := 0
	for i := 0; i < len(nums); i++ {
		//
		if i > k {
			return false
		}
		// 把【能跳到最远的距离】不断更新
		k = Max(k, i+nums[i])
	}
	return true
}

/*
	其实跳几步无所谓，关键在于可跳的覆盖范围！
	不一定非要明确一次究竟跳几步，每次取最大的跳跃步数，这个就是可以跳跃的覆盖范围。
	这个范围内，别管是怎么跳的，反正一定可以跳过来。
	那么这个问题就转化为跳跃覆盖范围究竟可不可以覆盖到终点！
	每次移动取最大跳跃步数（得到最大的覆盖范围），每移动一个单位，就更新最大覆盖范围。
	贪心算法局部最优解：每次取最大跳跃步数（取最大覆盖范围），整体最优解：最后得到整体最大覆盖范围，看是否能到终点。


	i每次移动只能在cover的范围内移动，每移动一个元素，cover得到该元素数值（新的覆盖范围）的补充，让i继续移动下去。
	而cover每次只取 max(该元素数值补充后的范围, cover本身范围)。
	如果cover大于等于了终点下标，直接return true就可以了。
*/
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	// 能跳的最远距离
	cover := 0
	// 注意这里是小于等于cover
	for i := 0; i <= cover; i++ {
		cover = MaxInt(i+nums[i], cover)
		// 说明可以覆盖到终点了
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}