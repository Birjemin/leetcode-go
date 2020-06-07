## 问题
Design a data structure that supports the following two operations:
```
void addWord(word)
bool search(word)
```
search(word) can search a literal word or a regular expression string containing only letters a-z or .. A . means it can represent any one letter.

Example:
```
addWord("bad")
addWord("dad")
addWord("mad")
search("pad") -> false
search("bad") -> true
search(".ad") -> true
search("b..") -> true
```

Note:
You may assume that all words are consist of lowercase letters a-z.

## 分析
设计一个支持以下两种操作的数据结构：
```
void addWord(word)
bool search(word)
```
search(word) 可以搜索文字或正则表达式字符串，字符串只包含字母 . 或 a-z 。 . 可以表示任何一个字母。

208的思想，速度和208一样，都不快

## 最高分
```golang
type WordDictionary struct {
	root *Node
}

type Node struct {
	isEnd    bool
	char     rune
	children [26]*Node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	root := Node{isEnd: false, char: 0, children: [26]*Node{}}
	return WordDictionary{
		root: &root,
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &Node{isEnd: false, char: c, children: [26]*Node{}}
		}
		node = node.children[c-'a']
	}
	node.isEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return match(word, 0, this.root)
}

func match(word string, index int, node *Node) bool {
	if index == len(word) {
		return node.isEnd
	}
	if word[index] == '.' {
		for _, child := range node.children {
			if child != nil && match(word, index+1, child) {
				return true
			}
		}
		return false
	} else {
		next := node.children[word[index]-'a']
		return next != nil && match(word, index+1, next)
	}
}
```

## 实现
208的思想，bfs算法即可
```golang
type WordDictionary struct {
	node  bool
	words [26]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	length, t := len(word), this

	for i := 0; i < length; i++ {
		idx := word[i] - 'a'
		if t.words[idx] == nil {
			t.words[idx] = &WordDictionary{}
		}
		t = t.words[idx]
	}

	// end flag
	t.node = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
    return this.find(this.words, word, 0, len(word))
}

func (this *WordDictionary) find(tries [26]*WordDictionary, word string, i, length int) bool {
	// the end of word
	if i == length {
		if this.node {
			return true
		}
		return false
	}

	var queue [26]*WordDictionary
	if word[i] != '.' {
		idx := word[i] - 'a'
		if tries[idx] == nil {
			return false
		}
		queue[idx] = tries[idx]
	} else {
		queue = tries
	}

	// next letter
	i++
	for _, v := range queue {
		if v != nil && v.find(v.words, word, i, length) {
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
```

## 改进
208的思想
```golang
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

```

## 反思

## 参考