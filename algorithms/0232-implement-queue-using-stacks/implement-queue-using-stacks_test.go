package implement_queue_using_stacks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ast := assert.New(t)

	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	ast.Equal(1, queue.Peek(), "输入:%v")
	ast.Equal(1, queue.Pop(), "输入:%v")
	ast.Equal(2, queue.Pop(), "输入:%v")
	ast.Equal(3, queue.Pop(), "输入:%v")
	ast.Equal(true, queue.Empty(), "输入:%v")
}
