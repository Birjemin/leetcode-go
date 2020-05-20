## 问题
Implement the following operations of a queue using stacks.

- push(x) -- Push element x to the back of queue.
- pop() -- Removes the element from in front of queue.
- peek() -- Get the front element.
- empty() -- Return whether the queue is empty.

Example:
```
MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);  
queue.peek();  // returns 1
queue.pop();   // returns 1
queue.empty(); // returns false
```

Notes:

- You must use only standard operations of a stack -- which means only push to top, peek/pop from top, size, and is empty operations are valid.
- Depending on your language, stack may not be supported natively. You may simulate a stack by using a list or deque (double-ended queue), as long as you use only standard operations of a stack.
- You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).

## 分析
使用栈实现队列`container/list`(不可以使用queue实现哦~)

225题目对应

## 最高分
```golang
type MyQueue struct {
	DataStack Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{DataStack: Stack{[]int{}}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.DataStack.push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	reverse := Stack{}
	for !this.Empty() {
		reverse.push(this.DataStack.pop())
	}
	x := reverse.pop()
	for reverse.len()!=0 {
		this.DataStack.push(reverse.pop())
	}
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	reverse := Stack{}
	for !this.Empty() {
		reverse.push(this.DataStack.pop())
	}
	x := reverse.peek()
	for reverse.len()!=0 {
		this.DataStack.push(reverse.pop())
	}
	return x
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.DataStack.len() == 0
}

type Stack struct {
	Data []int
}

func (receiver *Stack) pop() int {
	i := receiver.Data[len(receiver.Data)-1]
	receiver.Data = receiver.Data[:len(receiver.Data)-1]
	return i
}

func (receiver *Stack) push(i int) {
	receiver.Data = append(receiver.Data, i)
}

func (receiver *Stack) peek() int {
	return receiver.Data[len(receiver.Data)-1]
}

func (receiver *Stack) len() int {
	return len(receiver.Data)
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
```

## 实现
两个栈实现queue
```golang
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

```

## 改进
```golang

```

## 反思

## 参考