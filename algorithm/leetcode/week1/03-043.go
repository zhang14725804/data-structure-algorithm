/*
	模拟手算乘法的过程
	todo：被byte，string，int类型问题卡住了😅
*/
func multiply(num1 string, num2 string) string {
	l1:=len(num1)
	l2:=len(num2)

	product:=make([]byte,l1+l2)

	// 注意
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			product[l1-i-1+l2-j-1] += (num1[i]-'0')*(num2[j]-'0')
		}
	}
	var t byte
	for i:=0;i<len(product);i++{
		// 
		x:=product[i]
		t+=x
		x = t%10
		t/=10
	}
	k := len(product)-1
	for product[k]=='0' && k>0{
		k--
	}
	// 最后被卡在这里😅
	var res string
	for i := k; i >=0; i-- {
		res+=string(product[i])
	}
	return res
}