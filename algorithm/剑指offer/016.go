func replaceSpaces(str string) string {
	res:=""
	// s类型是int32
	for _,s := range str{
		if string(s) == " "{
			res+="%20"
		}else{
			res+= string(s)
		}
	}
	return res
}