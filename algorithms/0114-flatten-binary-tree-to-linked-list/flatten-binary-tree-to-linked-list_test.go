package flatten_binary_tree_to_linked_list

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one *TreeNode
    two int
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
        {
            p: param{
                one: makeTreeNode([]int{1, 2, 5, 3, 4, 0, 6}),
                two: 22,
            },
            a: result{
                one: makeTreeNode([]int{1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6}),
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        flatten(p.one)
        ast.Equal(a.one, p.one, "输入:%v", q)
    }
}
