## 问题
Implement the following operations of a stack using queues.

- push(x) -- Push element x onto stack.
- pop() -- Removes the element on top of the stack.
- top() -- Get the top element.
- empty() -- Return whether the stack is empty.

Example:
```
MyStack stack = new MyStack();

stack.push(1);
stack.push(2);  
stack.top();   // returns 2
stack.pop();   // returns 2
stack.empty(); // returns false
```

Notes:

- You must use only standard operations of a queue -- which means only push to back, peek/pop from front, size, and is empty operations are valid.
- Depending on your language, queue may not be supported natively. You may simulate a queue by using a list or deque (double-ended queue), as long as you use only standard operations of a queue.
- You may assume that all operations are valid (for example, no pop or top operations will be called on an empty stack).

## 分析
用队列实现栈(关键点是使用队列实现)

## 最高分
```golang
type MyStack struct {
	a, b *Queue
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		a: NewQueue(),
		b: NewQueue(),
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	// queue a is empty
	if this.a.Len() == 0 {
		this.a, this.b = this.b, this.a
	}
	this.a.Push(x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.a.Len() == 0 {
		this.a, this.b = this.b, this.a
	}

	for this.a.Len() > 1 {
		this.b.Push(this.a.Pop())
	}
	return this.a.Pop()
}


/** Get the top element. */
func (this *MyStack) Top() int {
	ret := this.Pop()
	this.Push(ret)
	return ret
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.a.Len() == 0 {
		this.a, this.b = this.b, this.a
	}

	return this.a.Len() == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func NewQueue() *Queue {
	return &Queue{}
}

type Queue struct {
	nums []int
}

func (q *Queue) Push(n int) {
	q.nums = append(q.nums, n)
}

func (q *Queue) Pop() int {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

func (q *Queue) Len() int {
	return len(q.nums)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
```

## 实现
不符合要求，不是队列
```golang
type MyStack struct {
	stack []int
	length int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.stack = append(this.stack, x)
	this.length++
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.length == 0 {
		return 0
	}
	ret := this.stack[this.length-1]
	this.stack = this.stack[:this.length-1]
	this.length--
	return ret
}


/** Get the top element. */
func (this *MyStack) Top() int {
	if this.length == 0 {
		return 0
	}
	return this.stack[this.length-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.length == 0
}
```

## 改进
使用两个队列来模拟栈，其实就是两个队列来回传递，得到最后一个值
```golang
type MyStack struct {
	a, b *Queue
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		a: NewQueue(),
		b: NewQueue(),
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	// queue a is empty
	if this.a.Len() == 0 {
		this.a, this.b = this.b, this.a
	}
	this.a.Push(x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.a.Len() == 0 {
		this.a, this.b = this.b, this.a
	}

	for this.a.Len() > 1 {
		this.b.Push(this.a.Pop())
	}
	return this.a.Pop()
}


/** Get the top element. */
func (this *MyStack) Top() int {
	ret := this.Pop()
	this.Push(ret)
	return ret
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.a.Len() == 0 {
		this.a, this.b = this.b, this.a
	}

	return this.a.Len() == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func NewQueue() *Queue {
	return &Queue{}
}

type Queue struct {
	nums []int
}

func (q *Queue) Push(n int) {
	q.nums = append(q.nums, n)
}

func (q *Queue) Pop() int {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

func (q *Queue) Len() int {
	return len(q.nums)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
```

## 反思

## 参考