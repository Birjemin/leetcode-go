package min_stack

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
