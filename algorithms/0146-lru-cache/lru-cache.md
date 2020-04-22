## 问题
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

The cache is initialized with a positive capacity.

Follow up:
Could you do both operations in O(1) time complexity?

Example:
```
LRUCache cache = new LRUCache( 2 /* capacity */ );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
```

## 分析
使用队列实现LRU算法(双向链表~)

## 最高分
```golang
func Constructor(capacity int) LRUCache {
	m := make(map[int]*CacheNode)
	c := LRUCache{Cap: capacity, Map: m}
	return c
}

func (this *LRUCache) Get(key int) int {
	found, ok := this.Map[key]
	if !ok {
		return -1
	}
	if this.Head == found {
		return found.Val
	}
	if this.Tail == found {
		this.Tail = found.Prev
	}
	// move found to head
	if found.Next != nil {
		found.Next.Prev = found.Prev
	}
	if found.Prev != nil {
		found.Prev.Next = found.Next
	}
	this.Head.Prev, found.Nex = found, this.Head
	this.Head = found
	return found.Val
}

func (this *LRUCache) Put(key int, value int) {
	found, ok := this.Map[key]
	if ok {
		found.Val = value
		_ = this.Get(found.Key) // to move found node to head
		return
	}

	// add to head
	n := &CacheNode{Key: key, Val: value}

	if len(this.Map) == 0 {
		this.Tail = n
	} else {
		this.Head.Prev, n.Next = n, this.Head
	}
	this.Map[key], this.Head = n, n
	if this.Cap == this.Len {
		delete(this.Map, this.Tail.Key)
		this.Len, this.Tail = this.Len+1, this.Tail.Prev
	}
	this.Len++
}

type CacheNode struct {
	Next *CacheNode
	Prev *CacheNode
	Key  int
	Val  int
}

type LRUCache struct {
	Cap  int
	Len  int
	Head *CacheNode
	Tail *CacheNode
	Map  map[int]*CacheNode
}
```

## 实现
不会，翻阅答案才知道使用双向链表
```golang
type LRUCache struct {
    capacity int
    data     map[int]*list.Element
    queue    *list.List
}

type Item struct {
    key   int
    value int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity: capacity,
        data:     make(map[int]*list.Element, capacity),
        queue:    list.New(),
    }
}

func (this *LRUCache) Get(key int) int {
    // found it
    if l, ok := this.data[key]; ok {
        this.queue.MoveToFront(l)
        return l.Value.(Item).value
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {

    // find it
    if l, ok := this.data[key]; ok {
        l.Value = Item{key: key, value: value}
        this.queue.MoveToFront(l)
        return
    }
    // full just remove the last element
    if this.queue.Len() == this.capacity {
        tmp := this.queue.Back()
        delete(this.data, tmp.Value.(Item).key)
        this.queue.Remove(tmp)
    }

    this.data[key] = this.queue.PushFront(Item{key: key, value: value})
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
```

## 改进
```golang

```

## 反思

## 参考