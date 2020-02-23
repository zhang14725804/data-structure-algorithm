/*
	6. ZigZag Conversion

	等差数列（首项n-1，公差2(n-1)的等差数列）

	不太好理解
*/
func convert(si string, n int) string {
    if n==1{
		return si
	}
	s := []rune(si)
	// 
	var res []rune
	for i:=0;i<n;i++{
		
		if i==0 || i==n-1{
			for j:=i;j<len(s);j+=2*(n-1){
				res =append(res,s[j])
			}
		}else{
			// 中间的部分(注意这个循环条件写法)
			for j,k:=i,2*(n-1)-i;j<len(s) || k<len(s);j,k=j+2*(n-1),k+2*(n-1){
				
				if j<len(s){
                    res =append(res,s[j])
				}
				if k<len(s){
                    res =append(res,s[k])
				}
			}
		}
	}
	return string(res)
}