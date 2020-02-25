/*
	208. Implement Trie (Prefix Tree)

	todos：：不懂这个题
*/
type Trie struct {
	end int
	val byte
	sons [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	node := this
	for i:=0; i<len(word);i++{
		idx:=word[i] - 'a'
		if node.sons[idx] == nil{
			node.sons[idx] = &Trie{val:word[i]}
		}
		// 
		node = node.sons[idx]
	}
	node.end++
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node:=this
	for i:=0;i<len(word);i++{
		idx:=word[i]-'a'
		if node.sons[idx] == nil{
			return false
		}
		node = node.sons[idx]
	}
	if node.end>0{
		return true
	}
	return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node:=this
	for i:=0;i<len(prefix);i++{
		idx:=prefix[i]-'a'
		if node.sons[idx] == nil{
			return false
		}
		node = node.sons[idx]
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */