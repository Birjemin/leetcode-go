package convert_sorted_array_to_binary_search_tree

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []int
}

type result struct {
    one *TreeNode
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
        // {
        //     p: param{
        //         one: []int{-10, -3, 1, 5, 9},
        //     },
        //     a: result{
        //         one: makeTreeNode([]int{1, -3, 9, -10, 0, 5}),
        //     },
        // },
        {
            p: param{
                one: []int{3, 5, 8},
            },
            a: result{
                one: makeTreeNode([]int{5, 3, 8}),
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, sortedArrayToBST(p.one), "输入:%v", q)
        ast.Equal(a.one, sortedArrayToBST1(p.one), "输入:%v", q)
    }
}
