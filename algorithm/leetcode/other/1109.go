/*
	差分数组 (😅😅😅 question)
	注意到一个预订记录实际上代表了一个区间的增量。我们的任务是将这些增量叠加得到答案。
	因此，我们可以使用差分解决本题。
*/
func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	for _, booking := range bookings {
		// 日期从 1 开始
		nums[booking[0]-1] += booking[2]
		if booking[1] < n {
			nums[booking[1]] -= booking[2]
		}
	}
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

/*
	我的暴力做法 😅
*/
func corpFlightBookings(bookings [][]int, n int) []int {
	count := make([]int, n)
	for i := 0; i < len(bookings); i++ {
		for j := bookings[i][0]; j <= bookings[i][1]; j++ {
			count[j-1] += bookings[i][2]
		}
	}
	return count
}