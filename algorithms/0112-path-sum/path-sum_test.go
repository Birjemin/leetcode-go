package path_sum

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one *TreeNode
    two int
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
                one: makeTreeNode([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 1}),
                two: 22,
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{1, 2, 0, 3, 0, 4, 0, 5}),
                two: 6,
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{1, 2}),
                two: 2,
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{1}),
                two: 1,
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{-2, 0, -3}),
                two: -5,
            },
            a: result{
                one: true,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, hasPathSum(p.one, p.two), "输入:%v", q)
        ast.Equal(a.one, hasPathSum1(p.one, p.two), "输入:%v", q)
    }
}
