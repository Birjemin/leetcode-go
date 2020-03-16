package linked_list_cycle

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
    one bool
}

type question struct {
    p param
    a result
}

func makeListNode(params []int, pos int) *ListNode {
    if len(params) == 0 {
        return nil
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
                one: makeListNode([]int{3, 2, 0, -4}, 1),
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: makeListNode([]int{1, 2}, 0),
            },
            a: result{
                one: true,
            },
        },
        {
            p: param{
                one: makeListNode([]int{1}, -1),
            },
            a: result{
                one: false,
            },
        },
    }

    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, hasCycle(p.one), "输入:%v", p)
    }
}
