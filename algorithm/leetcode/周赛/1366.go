/*
	todo：不懂
*/
func rankTeams(votes []string) string {
	n:=len(votes[0])
	// 二维数组
	ranks:=make([][]int,26)
	for i:=0;i<26 ;i++  {
		ranks[i] = make([]int,n+1)
	}
	for i:=0;i<26 ;i++  {
		// todo：这不懂
		ranks[i][n] = -i
	}

	// 每个评委的平分
	for _,vote:=range votes{
		for i:=0;i<n ;i++  {
			ranks[vote[i]-'A'][i] ++
		}
	}
	fmt.Println(ranks)
	res:=""
	// 什么骚操作 不懂了
	// 先排序
	//for i:=0;i<n ;i++  {
	//	res += -ranks[i][n]+'A'
	//}
	//res
	return res
}