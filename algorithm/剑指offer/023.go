/*
	暴力搜索，枚举所有可能情况。
	枚举起点，枚举方向，走到不能走为止
	todo：经典问题
*/
var arr1 = [][]byte{
	{'O', 'L', 'S', 'J', 'C', 'W', 'P', 'L'}, 
	{'U', 'R', 'I', 'P', 'T', 'D', 'E', 'I'},
	{'T', 'L', 'Z', 'E', 'G', 'B', 'V', 'X'}, 
	{'O', 'C', 'E', 'C', 'B', 'X', 'L', 'C'}, 
	{'I', 'L', 'R', 'X', 'A', 'J', 'C', 'M'}, 
	{'I', 'N', 'X', 'M', 'P', 'U', 'D', 'C'},
}
var target = "ITSQNSALO"
arr2 = [][]byte{{'a'}} 
var target2 = "a"
func hasPath(matrix [][]byte, str string) bool {
	// 遍历每个字符
    for i:=0; i<len(matrix); i++{
		for j:=0; j<len(matrix[i]); j++{
			if dfs(matrix,str,0,i,j){
				return true
			}
		}
	}
	return false
}
// u 当前字符的下边（当前枚举到str的第几个字符）
func dfs(matrix [][]byte, str string, u,x,y int) bool{
	if u == len(str){
		return true
	}
	if matrix[x][y] != str[u]{
		return false
	}
	// 四个方向(上右下左)，和坐标中的xy相反
	dx:=[]int{-1,0,1,0}
	dy:=[]int{0,1,0,-1}
	// 占位符，标记已经走过
	t := matrix[x][y]
	matrix[x][y] = '*'

	for i:=0; i<4; i++{
		a := x+dx[i]
		b := y+dy[i]
		// todo：判断边界；居然栽倒这里了😅
		if a>=0 && a<len(matrix) && b>=0 && b<len(matrix[a]){
			if dfs(matrix,str,u+1,a,b){
				return true
			}
		}
	}
	matrix[x][y] = t
	return false
}