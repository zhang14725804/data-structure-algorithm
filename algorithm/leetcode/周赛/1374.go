func generateTheString(n int) string {
	res:=""
	if n%2 == 0{
		res+="b"
		n--
	}
	for n>0{
		res+="a"
		n--
	}
	return res
}