/*
	arr := []int{3, 5, 1}
	arr := []int{0, 5, 3, 4, 6, 9}
*/

func canMakeArithmeticProgression(arr []int) bool {
	nLen := len(arr)
	sum := 0
	min := arr[0]
	max := arr[0]
	for i := 0; i < nLen; i++ {
		sum += arr[i]
		min = Min(min, arr[i])
		max = Max(max, arr[i])
	}
	if float64(sum) != float64(nLen*(min+max))/2 {
		return false
	}

	// TODO 如何判断公差是否是小数
	step := max - min/nLen - 1
	fstep := float64(max-min) / float64(nLen-1)
	isxiaoshu := false
	if fstep == float64(int64(fstep)) {
		isxiaoshu = true
	}
	for i := 1; i < nLen; i++ {
		if isxiaoshu && float64(arr[i]-arr[i-1])/fstep != 0 || (!isxiaoshu && (arr[i]-arr[i-1])/step != 0) {
			return false
		}
	}
	return true
}