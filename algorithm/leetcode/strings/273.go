/*
	273. Integer to English Words

	英文数字读法，三个三个分一组
*/
/*	
	syntax error: unexpected {, expecting expression (solution.go)
	var small [2]string{"1","2"}
*/ 
import "strings"
var small = [20]string {"Zero","One","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten",
					"Eleven","Twelve","Thirteen","Fourteen","Fifteen","Sixteen","Seventeen","Eighteen","Nineteen"}
var decade = [10]string{"","","Twenty","Thirty","Forty","Fifty","Sixty","Seventy","Eighty","Ninety"}
var big = [4]string{"Billion","Million","Thousand",""}
func numberToWords(num int) string {
	if num == 0 {
		return small[0]
	}
	res:=""
	for i,j:=1000000000,0;i>0;i,j=i/1000,j+1{
		if num >= i{
			res+=getPart(num/i) + big[j] + " "
			num %=i
		}
	}
	return strings.TrimSpace(res)
}
func getPart(num int) string{
	res:=""
	if num>=100{
		res+=small[num/100] + " Hundred "
		num%=100
	}

	if num==0{
		return res
	}

	if num>=20{
		res+=decade[num/10] + " "
		num%=10
	}

	if num==0{
		return res
	}

	res+=small[num] + " "

	return res
}