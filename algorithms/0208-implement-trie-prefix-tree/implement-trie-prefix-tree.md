## 问题
Implement a trie with insert, search, and startsWith methods.

Example:
```
Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");   
trie.search("app");     // returns true
```

Note:

- You may assume that all inputs are consist of lowercase letters a-z.
- All inputs are guaranteed to be non-empty strings.

## 分析
实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

Trie的概念可以了解一下，其实就是一个树的结构。

## 最高分
速度都不快~~~
```golang
type Trie struct {
    m map[byte]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie {
        m: make(map[byte]*Trie),
    }
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    current := this
    for i := 0; i < len(word); i++ {
        w := word[i]
        if v, ok := current.m[w]; ok {
            current = v
        } else {
            current.m[w] = &Trie{m: make(map[byte]*Trie)}
            current = current.m[w]
        }
    }
    // represents the end of the string
    current.m['0'] = &Trie{m: make(map[byte]*Trie)}
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    current := this
    for i := 0; i < len(word); i++ {
        v, ok := current.m[word[i]]
        if !ok {
            return false
        }
        current = v
    }
    // if it is not terminated, this string isn't included in map
    if _, ok := current.m['0']; !ok {
        return false
    }
    return true
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    current := this
    for i := 0; i < len(prefix); i++ {
        v, ok := current.m[prefix[i]]
        if !ok {
            return false
        }
        current = v
    }
    return true
}

```

## 实现
最基本的实现，insert时分配空间
```golang
type Trie struct {
	node  bool
	tries map[uint8]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		tries: make(map[uint8]*Trie),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if word == "" {
		this.node = true
		return
	}
	if _, ok := this.tries[word[0]]; !ok {
		this.tries[word[0]] = &Trie{
			tries: make(map[uint8]*Trie),
		}
	}
	this.tries[word[0]].Insert(word[1:])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if word == "" {
		if this.node {
			return true
		}
		return false
	}
	if trie, ok := this.tries[word[0]]; ok {
		return trie.Search(word[1:])
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return true
	}
	if trie, ok := this.tries[prefix[0]]; ok {
		return trie.StartsWith(prefix[1:])
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
```

## 改进
使用数组，提前分配空间
```golang
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

```

## 反思

## 参考