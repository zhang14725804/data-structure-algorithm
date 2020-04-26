/*
	行列递增的二维数组中的查找
*/
func searchArray(array [][]int, target int) bool {
    if len(array) == 0 || len(array[0]) == 0 {
		return false
	}
	// 右上角开始往下找
	i,j := 0,len(array[0])-1
	// 从右上角走到左下角
	for i < len(array) && j >= 0{
		if array[i][j] == target {
			return true
		}
		// 技巧，每次跳过一行或者一列
		if array[i][j] > target{
			j--
		}else{
			i++
		}
	}
	return false
}