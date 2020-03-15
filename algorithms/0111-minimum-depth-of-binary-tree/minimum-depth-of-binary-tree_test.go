package minimum_depth_of_binary_tree

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one *TreeNode
}

type result struct {
    one int
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
                one: makeTreeNode([]int{}),
            },
            a: result{
                one: 0,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{1}),
            },
            a: result{
                one: 1,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{1, 2, 3}),
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{1, 2}),
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{1, 2, 3, 4, 0, 0, 5}),
            },
            a: result{
                one: 3,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{3, 9, 20, 0, 0, 15, 7}),
            },
            a: result{
                one: 2,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{1, 2, 3, 4, 5}),
            },
            a: result{
                one: 2,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, minDepth(p.one), "输入:%v", q)
        ast.Equal(a.one, minDepth1(p.one), "输入:%v", q)
    }
}
