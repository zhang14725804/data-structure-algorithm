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
		// 排除”+“和”.“，生成name
		for _,c :=range(email[:at]){
			if c == '+' {
				break
			}else if c!='.'{
				name += string(c)
			}
		}

		domain:=email[at+1:]
		// 每个email存入hash表
		emain :=name + "@" + string(domain) 
		hash[emain] = 1
	}
	// 统计hash表长度（办法好像有点蠢）
	res:=0
	for _,_=range hash{
		res+=1
	}
	return res
}
// rune中查找元素返回元素下标
func find(arr []rune, char rune) int{
	for i:=0;i<len(arr);i++{
		// 注意这里
		if  arr[i] == char{
			return i
		}
	}
	return 0
}