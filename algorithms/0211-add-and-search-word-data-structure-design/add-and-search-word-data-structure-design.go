package add_and_search_word_data_structure_design

type WordDictionary struct {
	node  bool
	words map[uint8]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		words: make(map[uint8]*WordDictionary),
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	if word == "" {
		this.node = true
		return
	}
	if _, ok := this.words[word[0]]; !ok {
		this.words[word[0]] = &WordDictionary{
			words: make(map[uint8]*WordDictionary),
		}
	}
	this.words[word[0]].AddWord(word[1:])
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	// the end of word
	if word == "" {
		if this.node {
			return true
		}
		return false
	}

	var queue map[uint8]*WordDictionary
	if word[0] != '.' {
		if _, ok := this.words[word[0]]; !ok {
			return false
		} else {
			queue = make(map[uint8]*WordDictionary, 1)
			queue[word[0]] = this.words[word[0]]
		}
	} else {
		queue = this.words
	}

	// next letter
	for _, v := range queue {
		if v != nil && v.Search(word[1:]) {
			return true
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
