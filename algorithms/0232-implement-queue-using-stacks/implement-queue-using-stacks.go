package implement_queue_using_stacks

import (
	"container/list"
)

type MyQueue struct {
	// PushFront()
	// Front() and Remove()
	stack1, stack2 *list.List
	peek           int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack1: list.New(),
		stack2: list.New(),
		peek:   0,
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	if this.stack1.Len() == 0 {
		this.peek = x
	}
	this.stack1.PushFront(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {

	var ret *list.Element
	if this.stack1.Len() == 0 {
		panic("err")
	}

	this.peek = 0

	for i := this.stack1.Len(); i > 0; i-- {
		ret = this.stack1.Front()
		this.stack1.Remove(ret)
		if i != 1 {
			this.peek = ret.Value.(int)
			this.stack2.PushFront(this.peek)
		}
	}

	// traverse stack
	for this.stack2.Len() > 0 {
		tmp := this.stack2.Front()
		this.stack2.Remove(tmp)
		this.stack1.PushFront(tmp.Value.(int))
	}

	return ret.Value.(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.peek
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stack1.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
