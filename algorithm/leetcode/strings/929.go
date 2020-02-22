/*
	929. Unique Email Addresses

	golang没有set，所以用map替代set
*/
func numUniqueEmails(emails []string) int {
	hash:= make(map[string]int,0)
	for _,e := range emails{
		email:=[]rune(e)
		at:= find(email,'@')
		var name string
		for _,c :=range(email[:at]){
			if c == '+' {
				break
			}else if c!='.'{
				name += string(c)
			}
		}

		domain:=email[at+1:]
		emain :=name + "@" + string(domain) 
		hash[emain] = 1
	}
	res:=0
	for _,_=range hash{
		res+=1
	}
	return res
}
func find(arr []rune, char rune) int{
	for i:=0;i<len(arr);i++{
		// 注意这里
		a := arr[i]
		if  a== char{
			return i
		}
	}
	return 0
}