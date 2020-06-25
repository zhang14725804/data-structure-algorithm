/*
	6. ZigZag Conversion

	找规律；第一行和最后一行是公差2(n-1)的等差数列，中间的行是交错的等差数列

*/
func convert(si string, n int) string {
    if n==1{
		return si
	}
	s := []rune(si)
	// 枚举行
	var res []rune
	for i:=0;i<n;i++{
		// 第一行和最后一行（公差2(n-1)的等差数列）
		if i==0 || i==n-1{
			for j:=i; j<len(s); j+=2*(n-1){
				res =append(res,s[j])
			}
		}else{
			// 中间的部分(两个等差数列交错) j,k表示第一个第二个等差数列
			for j,k:=i,2*(n-1)-i; j<len(s) || k<len(s); j,k=j+2*(n-1),k+2*(n-1){
				
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