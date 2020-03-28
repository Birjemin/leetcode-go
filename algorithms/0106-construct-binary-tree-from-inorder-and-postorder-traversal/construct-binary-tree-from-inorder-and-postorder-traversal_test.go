package construct_binary_tree_from_inorder_and_postorder_traversal

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []int
    two []int
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
                one: []int{9, 3, 15, 20, 7},
                two: []int{9, 15, 7, 20, 3},
            },
            a: result{
                one: makeTreeNode([]int{3, 9, 20, 0, 0, 15, 7}),
            },
        },
        {
            p: param{
                one: []int{4, 2, 5, 1, 6, 3},
                two: []int{4, 5, 2, 6, 3, 1},
            },
            a: result{
                one: makeTreeNode([]int{1, 2, 3, 4, 5, 6, 0}),
            },
        },
        {
            p: param{
                one: []int{},
                two: []int{},
            },
            a: result{
                one: makeTreeNode([]int{}),
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, buildTree(p.one, p.two), "输入:%v", q)
        ast.Equal(a.one, buildTree1(p.one, p.two), "输入:%v", q)
    }
}
