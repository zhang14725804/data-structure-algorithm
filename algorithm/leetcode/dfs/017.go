/*
	17. Letter Combinations of a Phone Number
*/
func letterCombinations(digits string) []string {
	chars:=[]string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	if len(digits) == 0{
        return []string{}
	}
	// todos::为什么这里要有个空字符串
    state:=[]string{""}
	for _,u:=range digits{
		now:=[]string{}
		for _,c:=range chars[u-'2']{
			for _,s:=range state{
				now=append(now,string(s)+string(c))
			}
		}
		state = now
	}
	return state
}