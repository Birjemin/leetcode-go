package implement_trie_prefix_tree

type Trie struct {
	node  bool
	tries [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {

	length, t := len(word), this

	for i := 0; i < length; i++ {
		idx := word[i] - 'a'
		if t.tries[idx] == nil {
			t.tries[idx] = &Trie{}
		}
		t = t.tries[idx]
	}

	// end flag
	t.node = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {

	length, t := len(word), this

	for i := 0; i < length; i++ {
		idx := word[i] - 'a'
		if t.tries[idx] == nil {
			return false
		}
		t = t.tries[idx]
	}

	if t.node {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {

	length, t := len(prefix), this

	for i := 0; i < length; i++ {
		idx := prefix[i] - 'a'
		if t.tries[idx] == nil {
			return false
		}
		t = t.tries[idx]
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
