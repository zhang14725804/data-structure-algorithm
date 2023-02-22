/*
	给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
	思路：
	用两个变量分别保存第一行第一列是否有0
	循环每个字符，如果是0，把当前字符的所在行和列的第一个数字置0
	循环完成之后，遍历第一行第一列，遇到0的，将该行该列置0
	最后检查最开始定义的两个变量，是否为0，来决定是否把第一行第一列置为0
*/
func setZeroes(m [][]int)  {
	firstRow:=false
	firstCol:=false
	// 检查第一列
	for i := 0; i < len(m)-1; i++ {
		if m[i][0] == 0{
			firstCol = true
		}
	}
	// 检查第一行
	for i := 0; i < len(m[0])-1; i++ {
		if m[0][i] == 0{
			firstRow = true
		}
	}
	// 用第一行第一列，标记改行该列是否有0
	for row := 0; row < len(m)-1; row++ {
		for col := 0; col < len(m[0])-1; col++ {
			if m[row][col] == 0{
				m[row][0] = 0
				m[0][col] = 0
			}
		}
	}
	/*
		todo：代码有问题
	*/
	for row := 0; row < len(m)-1; row++ {
		for col := 0; col < len(m[0])-1; col++ {
			if m[row][0] == 0 || m[0][col] == 0{
				m[row][col] = 0
			}
		}
	}
	if firstRow == true{
		for i := 0; i < len(m[0])-1; i++ {
			m[0][i] = 0
		}
	}
	if firstCol == true{
		for i := 0; i < len(m)-1; i++ {
			m[i][0] = 0
		}
	}
}