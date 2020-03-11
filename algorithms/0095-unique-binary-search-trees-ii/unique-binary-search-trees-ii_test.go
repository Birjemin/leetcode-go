package unique_binary_search_trees_ii

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one int
}

type result struct {
    one []*TreeNode
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
                one: 3,
            },
            a: result{
                one: []*TreeNode{
                    makeTreeNode([]int{1, 0, 2, 0, 3}),
                    makeTreeNode([]int{1, 0, 3, 2}),
                    makeTreeNode([]int{2, 1, 3}),
                    makeTreeNode([]int{3, 1, 0, 0, 2}),
                    makeTreeNode([]int{3, 2, 0, 1}),
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, generateTrees(p.one), "输入:%v", q)
    }
}
