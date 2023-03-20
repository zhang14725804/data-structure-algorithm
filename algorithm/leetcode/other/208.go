/*
	前缀树的用例：词频统计，前缀匹配，路由匹配
*/
type Trie struct {
	end  bool      // 判断是否以这个字母结尾
	val  byte      // 父层字母
	sons [26]*Trie // 子层出现的字母
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (node *Trie) Insert(word string) {
	// 从前往后遍历每一个字母
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		// 判断是否存在
		if node.sons[idx] == nil {
			node.sons[idx] = &Trie{val: word[i]}
		}
		// 走到下一个节点
		node = node.sons[idx]
	}
	node.end = true
}

/** Returns if the word is in the trie. */
func (node *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		// 当前子节点不存在
		if node.sons[idx] == nil {
			return false
		}
		// 走到下一个节点
		node = node.sons[idx]
	}
	// 判断是否以这个节点结尾
	return node.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (node *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		idx := prefix[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}
		// 走到下一个节点
		node = node.sons[idx]
	}
	return true
}