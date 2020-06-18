/*
	todo:难懂
*/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 || len(image[0])==0{
		return image
	}
	// 四个方向(纵轴下，横轴y)
	dx:=[4]int{-1,0,1,0}
	dy:=[4]int{0,1,0,-1}
	oldColor := image[sr][sc]
	// 
	if oldColor == newColor{
		return image
	}
	image[sr][sc] = newColor
	// 遍历四个方向
	for i:=0;i<4;i++{
		x:=sr+dx[i]
		y:=sc+dy[i]
		if x>=0 && x<len(image) && y>=0 && y<len(image[0]) && image[x][y] == oldColor{
			floodFill(image,x,y,newColor)
		}
	}
	return image
}