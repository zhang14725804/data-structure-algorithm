/*
	165. Compare Version Numbers

	reflect.TypeOf(x)判断类型

	todos:执行的时候可以通过，提交的时候遇到"01"和"1"的情况无法通过
*/
func compareVersion(version1 string, version2 string) int {
	s1:=[]rune(version1)
	s2:=[]rune(version2)

	i,j:=0,0
	for i<len(s1) || j<len(s2){
		x,y:=i,j

		for x < len(s1) && s1[x] != '.' {
			x = x+1
		}

		for y < len(s2) && s2[y] != '.' {
			y = y+1
		}

		var a int
		if i == x{
			a = 0
		}else{
			a,_ = strconv.Atoi(string(s1[i:i+1]))
		}
			
		var b int
		if j == y{
			b = 0
		}else{
			b,_ = strconv.Atoi(string(s2[j:j+1]))
		}

		if a>b{
			return 1
		}
		if a<b{
			return -1
		}
		i = x+1
		j = y+1
	}
	return 0
}