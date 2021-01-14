/*
	给定一个未排序的整数数组，找出最长连续序列的长度。
	要求算法的时间复杂度为 O(n)。
*/

// 思路，先排序，然后找
func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	res := 0
	j := 1
	for i := 0; i < len(nums); i++ {
		// 去重
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
		if i+1 < len(nums) && nums[i]+1 == nums[i+1] {
			j++
		} else {
			res = compare(res, j, true)
			j = 1
		}
	}
	return res
}

// hash的方法，每次从最小值（每一段的开始）开始枚举，枚举完之后删除
func longestConsecutive(nums []int) int {
	// 存入set中
	set:=NewSet()
	for _,n:=range nums{
		set.Insert(n)
	}

	res:=0
	for _,x:=range nums{
		// 每次从最小值（每一段的开始）开始枚举
		if set.Contains(x) && !set.Contains(x-1){
			y:=x
			// 枚举完之后删除
			set.Remove(x)
			for set.Contains(y+1){
				y++
				set.Remove(y)
			}
			res = MaxInt(res,y-x+1)
		}
	}
	return res
}