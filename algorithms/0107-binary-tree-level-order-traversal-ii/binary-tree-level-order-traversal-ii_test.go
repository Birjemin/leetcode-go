package binary_tree_level_order_traversal_ii

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one *TreeNode
}

type result struct {
    one [][]int
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
                one: makeTreeNode([]int{3, 9, 20, 0, 0, 15, 7}),
            },
            a: result{
                one: [][]int{
                    {15, 7},
                    {9, 20},
                    {3},
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, levelOrderBottom(p.one), "输入:%v", q)
        ast.Equal(a.one, levelOrderBottom1(p.one), "输入:%v", q)
    }
}
