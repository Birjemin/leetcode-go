package populating_next_right_pointersin_each_node

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one *Node
}

type result struct {
    one []int
}

type question struct {
    p param
    a result
}

func makeTreeNode(ints []int) *Node {
    n := len(ints)
    if n == 0 {
        return nil
    }
    root := &Node{
        Val: ints[0],
    }
    queue := make([]*Node, 1, n*2)
    queue[0] = root
    i := 1
    for i < n {
        node := queue[0]
        queue = queue[1:]
        if i < n && ints[i] != 0 {
            node.Left = &Node{Val: ints[i]}
            queue = append(queue, node.Left)
        }
        i++
        if i < n && ints[i] != 0 {
            node.Right = &Node{Val: ints[i]}
            queue = append(queue, node.Right)
        }
        i++
    }
    return root
}

func showTreeNode(tmp []*Node, ret *[]int) {
    if tmp == nil {
        return
    }
    var t []*Node
    var p *Node
    for _, v := range tmp {
        *ret = append(*ret, v.Val)
        p=v
        if v.Left != nil {
            t = append(t, v.Left, v.Right)
        }
    }
    if p.Next == nil {
        *ret = append(*ret, 0)
    }
    showTreeNode(t, ret)
}

func Test(t *testing.T) {
    ast := assert.New(t)
    qs := []question{
        {
            p: param{
                one: makeTreeNode([]int{1, 2, 3, 4, 5, 6, 7}),
            },
            a: result{
                one: []int{1, 0, 2, 3, 0, 4, 5, 6, 7, 0},
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        var ret []int
        showTreeNode([]*Node{connect1(p.one)}, &ret)
        ast.Equal(a.one, ret, "输入:%v", q)
    }
}
