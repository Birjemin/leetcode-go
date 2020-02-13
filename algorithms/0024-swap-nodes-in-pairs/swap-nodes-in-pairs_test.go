package swap_nodes_in_pairs

import (
    "fmt"
    "github.com/spf13/cast"
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one *ListNode
}

type result struct {
    one *ListNode
}

type question struct {
    p param
    a result
}

func makeListNode(params []int) *ListNode {
    if len(params) == 0 {
        return nil
    }
    res := &ListNode{
        Val: params[0],
    }
    temp := res
    for i, v := range params {
        if i != 0 {
            temp.Next = &ListNode{Val: v}
            temp = temp.Next
        }
    }
    return res
}

func showListNode(res *ListNode) {
    str := ""
    for res != nil {
        str += cast.ToString(res.Val) + " -> "
        res = res.Next
    }
    str += "\n"
    fmt.Print(str)
}

func Test(t *testing.T) {
    ast := assert.New(t)
    qs := []question{
        {
            p: param{
                one: makeListNode([]int{1, 2, 3, 4}),
            },
            a: result{
                one: makeListNode([]int{2, 1, 4, 3}),
            },
        },
        {
            p: param{
                one: makeListNode([]int{}),
            },
            a: result{
                one: makeListNode([]int{}),
            },
        },
        {
            p: param{
                one: makeListNode([]int{1, 2, 3}),
            },
            a: result{
                one: makeListNode([]int{2, 1, 3}),
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, swapPairs(p.one), "输入:%v", q)
    }
}
