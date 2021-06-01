/*
	把数组排成最小的数
	todo：思路清奇
*/
func printMinNumber(nums []int) string {
	BubbleSort(nums)
	res := ""
	for _, x := range nums {
		res += strconv.Itoa(x)
	}
	return res
}

func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if com(nums[j], nums[j+1]) {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 蠢 这么多 strconv.Itoa 😅 自定义比较规则 a<b <=> ab<ba
func com(a, b int) bool {
	res := strings.Compare(strconv.Itoa(a)+strconv.Itoa(b), strconv.Itoa(b)+strconv.Itoa(a))
	if res < 0 {
		return false
	}
	return true
}