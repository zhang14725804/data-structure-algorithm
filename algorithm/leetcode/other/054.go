/*
	给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

	思路：右下左上循环遍历
	0103 东信北邮 笔试题遇到
*/
func spiralOrder(matrix [][]int) []int {
    ans:=make([]int,0)
    if len(matrix)==0{
        return ans
    }

	x,y := 0,0
	// 0向右 1向下 2向左 3向上
	direct:=0
	// 上下左右边界
    top:=-1
    bottom:=len(matrix)
    left:=-1
    right:=len(matrix[0])

    for {
		// base case出口
        if len(ans) == len(matrix)*len(matrix[0]){
            return ans
        }
		// question： 注意 y 方向写在前边，x 方向写在后边
		ans = append(ans,matrix[y][x])
		// 右、下、左、上四个方向走
        switch direct{
            case 0: // 向右走
                if x+1 == right{
                    direct=1
                    y++
                    top++
                }else{
                    x++
                }
            case 1: // 向下走
                if y+1 == bottom{
                    direct=2
                    x--
                    right--
                }else{
                    y++
                }
            case 2: // 向左走
                if x-1 == left{
                    direct=3
                    y--
                    bottom--
                }else{
                    x--
                }
            case 3:
                if y-1==top{
                    direct=0
                    x++
                    left++
                }else{
                    y--
                }
        }     
    }
}
