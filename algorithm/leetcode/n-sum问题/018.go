/*
	给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
	答案中不可以包含重复的四元组。
*/

/*
	双指针法：

	使用四个指针(a<b<c<d)。固定最小的a和b在左边，c=b+1,d=_size-1 移动两个指针包夹求解。
	保存使得nums[a]+nums[b]+nums[c]+nums[d]==target的解。
	偏大时d左移，偏小时c右移。
	c和d相遇时，表示以当前的a和b为最小值的解已经全部求得。
	b++,进入下一轮循环b循环，当b循环结束后。
	a++，进入下一轮a循环。 即(a在最外层循环，里面嵌套b循环，再嵌套双指针c,d包夹求解)。

*/

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	// 排序先
	nums = sort(nums)

	var res [][]int
	for i := 0; i < n; i++ {
		// 去重
		if i == 0 || nums[i] != nums[i-1] {
			for j := i + 1; j < n-1; j++ {
				// 去重
				if j == i+1 || nums[j] != nums[j-1] {
					start, end := j+1, n-1
					// 移动指针😅
					for start < end {
						if nums[i]+nums[j]+nums[start]+nums[end] == target {
							res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})
							start++
							end--
							// 去重
							for start < end && nums[start] == nums[start-1] {
								start++
							}
							for start < end && nums[end] == nums[end+1] {
								end--
							}
						} else if nums[i]+nums[j]+nums[start]+nums[end] < target {
							start++
						} else {
							end--
						}
					}
				}
			}
		}
	}
	return res
}