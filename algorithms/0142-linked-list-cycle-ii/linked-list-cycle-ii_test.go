package linked_list_cycle_ii

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

func makeListNode(params []int, pos int) (*ListNode, *ListNode) {
    if len(params) == 0 {
        return nil, nil
    }
    var p *ListNode
    res := &ListNode{
        Val: params[0],
    }
    temp := res
    for i, v := range params {
        if i != 0 {
            temp.Next = &ListNode{Val: v}
            temp = temp.Next
        }
        if i == pos {
            p = temp
        }
    }
    temp.Next = p
    return res, p
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

    a, b := makeListNode([]int{3, 2, 0, -4}, 1)
    c, d := makeListNode([]int{1, 2}, 0)
    e, f := makeListNode([]int{1}, -1)

    qs := []question{
        {
            p: param{
                one: a,
            },
            a: result{
                one: b,
            },
        },
        {
            p: param{
                one: c,
            },
            a: result{
                one: d,
            },
        },
        {
            p: param{
                one: e,
            },
            a: result{
                one: f,
            },
        },
    }

    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, detectCycle(p.one), "输入:%v", p)
        ast.Equal(a.one, detectCycle1(p.one), "输入:%v", p)
    }
}
