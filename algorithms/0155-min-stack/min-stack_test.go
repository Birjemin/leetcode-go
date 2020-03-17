package min_stack

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test(t *testing.T) {
    ast := assert.New(t)
    minStack := Constructor()
    minStack.Push(-2)
    minStack.Push(0)
    minStack.Push(-3)
    ast.Equal(-3, minStack.GetMin(), "输入:%v", -3)
    minStack.Pop()
    ast.Equal(0, minStack.Top(), "输入:%v", 0)
    ast.Equal(-2, minStack.GetMin(), "输入:%v", -2)
}
