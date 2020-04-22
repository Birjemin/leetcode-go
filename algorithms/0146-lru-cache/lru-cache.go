package lru_cache

import (
    "container/list"
)

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
