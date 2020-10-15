/*
	给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

	双指针法：

		使用四个指针(a<b<c<d)。固定最小的a和b在左边，c=b+1,d=_size-1 移动两个指针包夹求解。
	保存使得nums[a]+nums[b]+nums[c]+nums[d]==target的解。偏大时d左移，偏小时c右移。c和d相
	遇时，表示以当前的a和b为最小值的解已经全部求得。b++,进入下一轮循环b循环，当b循环结束后。
	a++，进入下一轮a循环。 即(a在最外层循环，里面嵌套b循环，再嵌套双指针c,d包夹求解)。

*/
func fourSum(nums []int, target int) [][]int {
	nums = quickSort(nums)
	res := make([][]int, 0)
	n := len(nums)
	for i := 0; i < n-3; i++ {
		// 防止重复
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			for j := i + 1; j < n-2; j++ {
				// 防止重复
				if j == i+1 || nums[j] != nums[j-1] {
					start, end := j+1, n-1
					sum := target - nums[i] - nums[j]
					for start < end {
						if nums[start]+nums[end] == sum {
							res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})
							for start < end && nums[start] == nums[start+1] {
								start++
							}
							for start < end && nums[end] == nums[end-1] {
								end--
							}
							start++
							end--
						} else if nums[start]+nums[end] < sum {
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