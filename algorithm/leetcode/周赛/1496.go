/*
	给你一个字符串 path，其中 path[i] 的值可以是 'N'、'S'、'E' 或者 'W'，分别表示向北、向南、向东、向西移动一个单位。

	机器人从二维平面上的原点 (0, 0) 处开始出发，按 path 所指示的路径行走。

	如果路径在任何位置上出现相交的情况，也就是走到之前已经走过的位置，请返回 True ；否则，返回 False 。

	相交的点一定在path路径的各个端点，而不是线段中的某一点；用hash存储各个点，出现次数大于1的点就是相交的点
*/
func isPathCrossing(path string) bool {
	hash:=make(map[string]int,0)
	hash[string(0)+"_"+string(0)] = 1
	x,y:=0,0
	for _,c:=path{
		if c=='N'{
			x--
		}else if c =='S'{
			x++
		}else if c=='W'{
			y--
		}else{
			y++
		}
		s := string(x)+"_"+string(y)
		// todo：思路是这么个思路，代码有问题😅。"NESWW"\"WSWNNNNWWNWWN"测试不通过
		// golang这个string，坑
		if hash[s] > 0{
			return true
		}
		hash[s]++
	}
	return false
}