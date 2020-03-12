package validate_binary_search_tree

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one *TreeNode
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

    for i := 1; i < n; {
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
                one: makeTreeNode([]int{2, 1, 3}),
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{5, 1, 4, 0, 0, 3, 6}),
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{1, 1}),
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{10, 5, 15, 0, 0, 6, 20}),
            },
            a: result{
                one: false,
            },
        },
        {
            p: param{
                one: makeTreeNode([]int{0}),
            },
            a: result{
                one: true,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, isValidBST(p.one), "输入:%v", q)
        ast.Equal(a.one, isValidBST1(p.one), "输入:%v", q)
    }
}
