package implement_stack_using_queues

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