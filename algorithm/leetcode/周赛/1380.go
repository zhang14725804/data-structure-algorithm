/*

*/
func luckyNumbers (matrix [][]int) []int {
    n,m:=len(matrix),len(matrix[0])
	res:=make([]int,0)
	// 遍历所有元素
    for i:=0;i<n;i++{
        for j:=0;j<m;j++{
            lucky:=true
            // 枚举行，判断是否是最小
            for k:=0;k<m;k++{
                if matrix[i][j] > matrix[i][k]{
                    lucky = false
                    break
                }
            }
            // 枚举列，判断是否是最大
            if lucky{
                for k:=0;k<n;k++{
                    if matrix[i][j] < matrix[k][j]{
                        lucky = false
                        break
                    }
                }
                if lucky{
                    res= append(res,matrix[i][j])
                }
            }
        }
    }
    return res
}