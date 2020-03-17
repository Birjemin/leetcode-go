## 问题
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

- push(x) -- Push element x onto stack.
- pop() -- Removes the element on top of the stack.
- top() -- Get the top element.
- getMin() -- Retrieve the minimum element in the stack.

Example:
```
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
```

## 分析


## 最高分
```golang

```

## 实现
最基本的实现
```golang
type MinStack struct {
    params []int
    min    int
    top    int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}

func (this *MinStack) Push(x int) {
    this.top = x
    this.params = append(this.params, x)
}

func (this *MinStack) Pop() {
    length := len(this.params)
    this.params = this.params[:length-1]
    if length < 1 {
        return
    }
    this.top = this.params[length-1]
}

func (this *MinStack) Top() int {
    return this.top
}

func (this *MinStack) GetMin() int {
    min := this.params[0]
    for _, v := range this.params {
        if v < min {
            min = v
        }
    }
    return min
}
```

## 改进
必要时计算min，反之不计算，属于不稳定算法
```golang
type MinStack struct {
    params []int
    min    int
    top    int
    length int
    flag bool
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}

func (this *MinStack) Push(x int) {
    if this.length == 0 || this.min > x{
        this.min = x
    }
    this.length++
    this.top, this.params = x, append(this.params, x)
}

func (this *MinStack) Pop() {
    if this.params[this.length-1] == this.min {
        this.flag = true
    }
    this.params = this.params[:this.length-1]
    this.length--
    if this.length < 1 {
        return
    }
    this.top = this.params[this.length-1]
}

func (this *MinStack) Top() int {
    return this.top
}

func (this *MinStack) GetMin() int {
    if !this.flag {
        return this.min
    }
    this.min = this.params[0]
    for _, v := range this.params {
        if v < this.min {
            this.min = v
            this.flag = false
        }
    }
    return this.min
}
```

## 改进2
使用常数获取min的值，使用空间换区时间，稳定算法
- 记录每一次push时min的最小值即可，如果x>getMin，说明pop的时候getMin没有变，反之更新
```golang
type MinStack struct {
    data   []int
    min    []int
    length int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}

func (this *MinStack) Push(x int) {
    this.pushMin(x)
    this.data = append(this.data, x)
    this.length++
}

func (this *MinStack) Pop() {
    this.data, this.min = this.data[:this.length-1], this.min[:this.length-1]
    this.length--
}

func (this *MinStack) Top() int {
    return this.data[this.length-1]
}

func (this *MinStack) GetMin() int {
    return this.min[this.length-1]
}

func (this *MinStack) pushMin(x int) {
    var min int
    if this.length == 0 {
        min = x
    } else {
        min = this.GetMin()
        if min > x {
            min = x
        }
    }
    this.min = append(this.min, min)
}
```
## 反思

## 参考