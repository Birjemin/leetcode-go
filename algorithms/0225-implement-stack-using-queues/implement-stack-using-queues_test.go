package implement_stack_using_queues

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ast := assert.New(t)

	obj := Constructor()
	obj.Push(1)
	ast.Equal(1, obj.Pop(), "输入:%v")
	ast.Equal(true, obj.Empty(), "输入:%v")
}
