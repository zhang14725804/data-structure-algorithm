/*
	todo：整数划分的问题（思路不错😄）：拆分成，最多两个2和m个3（没有大于等于4的）
*/
func maxProductAfterCutting(l int) int {
    if l <= 3{
		return l-1
	}
	res := 1
	if l % 3 == 1{
		res *= 4
		n -= 4
	}
	if l % 3 == 2{
		res *= 2
		l -= 2
	}
	for l>0 {
		res *= 3
		l -= 3
	}
	return res
}