/*
	给定两个字符串 s 和 t，判断它们是否是同构的。
 */
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	hash := make(map[byte]byte,0)

	// s <==> t 相互映射
	for i:=0;i<len(s);i++{
		a,b := s[i],t[i]
        if val,ok:=hash[a];ok && val != b{
            return false
        }
		hash[a] = b
		
		if val,ok:=hash[b] ;ok && val!= a{
			return false
		}
		hash[b] = a
	}
	return true
}