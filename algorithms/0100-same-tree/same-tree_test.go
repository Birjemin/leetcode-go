package same_tree

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one *TreeNode
    two *TreeNode
}

type result struct {
    one bool
}

type question struct {
    p param
    a result
}

func makeTreeNode(ints []int) *TreeNode {
    n := len(ints)
    if n == 0 {
        return nil
    }
    root := &TreeNode{
        Val: ints[0],
    }
    queue := make([]*TreeNode, 1, n*2)
    queue[0] = root
    i := 1
    for i < n {
        node := queue[0]
        queue = queue[1:]
        if i < n && ints[i] != 0 {
            node.Left = &TreeNode{Val: ints[i]}
            queue = append(queue, node.Left)
        }
        i++
        if i < n && ints[i] != 0 {
            node.Right = &TreeNode{Val: ints[i]}
            queue = append(queue, node.Right)
        }
        i++
    }
    return root
}

func Test(t *testing.T) {
    ast := assert.New(t)
    qs := []question{
        {
            p: param{
                one: makeTreeNode([]int{1, 2, 3}),
                two: makeTreeNode([]int{1, 2, 3}),
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{1, 2}),
                two: makeTreeNode([]int{1, 0, 2}),
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{1, 2, 1}),
                two: makeTreeNode([]int{1, 1, 2}),
            },
            a: result{
                one: false,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, isSameTree(p.one, p.two), "输入:%v", q)
    }
}
