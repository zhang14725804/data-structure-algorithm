/*
	要从覆盖范围出发，不管怎么跳，覆盖范围内一定是可以跳到的，以最小的步数增加覆盖范围，覆盖范围一旦覆盖了终点，得到的就是最小步数！

	这里需要统计两个覆盖范围，当前这一步的最大覆盖和下一步最大覆盖。

	如果移动下标达到了当前这一步的最大覆盖最远距离了，还没有到终点的话，那么就必须再走一步来增加覆盖范围，直到覆盖范围覆盖了终点
*/

/*
	移动下标达到了当前覆盖的最远距离下标时，步数就要加一，来增加覆盖距离。最后的步数就是最少步数。

	这里还是有个特殊情况需要考虑，当移动下标达到了当前覆盖的最远距离下标时

	（1）如果当前覆盖最远距离下标不是是集合终点，步数就加一，还需要继续走。
	（2）如果当前覆盖最远距离下标就是是集合终点，步数不用加一，因为不能再往后走了。
*/
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	// 当前覆盖最远距离下标
	curDistance := 0
	// 下一步覆盖最远距离下标
	nextDistance := 0
	// 记录走的最大步数
	res := 0
	for i := 0; i < len(nums); i++ {
		// 更新下一步覆盖最远距离下标
		nextDistance = MaxInt(i+nums[i], nextDistance)
		// 遇到当前覆盖最远距离下标
		if i == curDistance {
			// 如果当前覆盖最远距离下标不是终点
			if curDistance != len(nums)-1 {
				// 更新操作
				res++
				curDistance = nextDistance
				// 下一步的覆盖范围已经可以达到终点，结束循环
				if nextDistance >= len(nums)-1 {
					break
				}
			} else {
				// 当前覆盖最远距离下标是集合终点，不用做ans++操作了，直接结束
				break
			}
		}
	}
	return res
}

/*
	版本2 😅😅😅

	移动下标只要遇到当前覆盖最远距离的下标，直接步数加一，不考虑是不是终点的情况。

	想要达到这样的效果，只要让移动下标，最大只能移动到nums.size - 2的地方就可以了。

	如果移动下标等于当前覆盖最大距离下标， 需要再走一步（即ans++），因为最后一步一定是可以到的终点。（题目假设总是可以到达数组

	如果移动下标不等于当前覆盖最大距离下标，说明当前覆盖最远距离就可以直接达到终点了，不需要再走一步。
*/
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	curDistance := 0
	nextDistance := 0
	res := 0
	// 😅😅😅 注意这里是小于nums.size() - 1，这是关键所在
	for i := 0; i < len(nums)-1; i++ {
		nextDistance = MaxInt(i+nums[i], nextDistance)
		if i == curDistance {
			res++
			curDistance = nextDistance
		}
	}
	return res
}