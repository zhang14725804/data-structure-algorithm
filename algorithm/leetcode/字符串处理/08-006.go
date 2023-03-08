/*
	1. 所有字符都按照公差N=2(n-1)方式排列
	2. 第一行和最后一行公差为N的等差数列
	3. 中间的部分是两个公车为N等差数列交错， j,k表示第一个第二个等差数列

	难点在于找规律😅😅😅
*/
func convert(si string, n int) string {
	if n == 1 {
		return si
	}

	var res []byte
	// 公差
	N := 2 * (n - 1)
	// 枚举行
	for i := 0; i < n; i++ {
		// 第一行和最后一行（公差2(n-1)的等差数列）
		if i == 0 || i == n-1 {
			for j := i; j < len(si); j += N {
				res = append(res, si[j])
			}
		} else {
			// 中间的部分(两个等差数列交错) j,k表示第一个第二个等差数列
			for j, k := i, N-i; j < len(si) || k < len(si); j, k = j+N, k+N {
				if j < len(si) {
					res = append(res, si[j])
				}
				if k < len(si) {
					res = append(res, si[k])
				}
			}
		}
	}
	return string(res)
}