/*
	某个性质可以把数组分为两部分，一部分满足，一部分不满足（二段性）
*/
func hIndex(citations []int) int {
	l,r := 0,len(citations)
	for l<r {
		mid := (l+r+1) >> 1
		// todo这里几个意思
		if citations[len(citations)-mid] >= mid{
			l = mid
		}else{
			r = mid-1
		}
	}
	return r
}