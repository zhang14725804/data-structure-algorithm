func firstBadVersion(n int) int {
	l,r := 0,n
	for l<r {
		mid := (l+r) >> 1
		if isBadVersion(mid) {
			r = mid
		}else{
			l = mid+1
		}
	}
	return r
}