/*
	todo:题意比较抽象😅，思路妙
*/
func numTimesAllBlue(light []int) int {
	// res答案，k最大值
	res,k := 0,0
	for i := 0; i < len(light); i++ {
		k = max(k,light[i])
		if k==i+1{
			res++
		}
	}
	return res
}
func max(x,y int) int{
	if x>y{
		return x
	}
	return y
}