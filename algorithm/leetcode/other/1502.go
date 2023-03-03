/*
	当前错误思路😅😅😅：
		1. 一次循环，找出最大值最小值，求和
		2. 判断是否符合等差数列求和公式算出来的值
		3. 算出公差，遍历数组，判断是否是公差整数倍

	arr := []int{3, 5, 1}
	arr := []int{0, 5, 3, 4, 6, 9}
	arr := []int{1, 2, 2, 5, 5}

	正确思路：
		1. 排序，并且求和，算出公差
		2. 遍历排序后的数组，检查是否符合公差数列特征
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
	step := (max - min) / (nLen - 1)
	if sum == 0 && step == 0 {
		return true
	}
	fstep := float64(max-min) / float64(nLen-1)
	isxiaoshu := false
	if fstep != float64(int64(step)) {
		isxiaoshu = true
	}
	for i := 1; i < nLen; i++ {
		if isxiaoshu && float64(arr[i]-arr[i-1])/fstep != 0 {
			return false
		}
		if !isxiaoshu && (arr[i]-arr[i-1])%step != 0 {
			return false
		}
	}
	return true
}