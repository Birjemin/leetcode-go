package intersection_of_two_linked_lists

import (
    "fmt"
    "github.com/spf13/cast"
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one *ListNode
    two *ListNode
}

type result struct {
    one *ListNode
}

type question struct {
    p param
    a result
}

func makeCircleListNode(params1, params2, params3 []int) (*ListNode, *ListNode, *ListNode) {

    l1, e1 := makeListNode(params1)
    l2, e2 := makeListNode(params2)
    l3, _ := makeListNode(params3)
    e1.Next, e2.Next = l3, l3
    return l1, l2, l3
}

func makeListNode(params []int) (*ListNode, *ListNode) {
    if len(params) == 0 {
        return nil, nil
    }
    head := &ListNode{
        Val: params[0],
    }
    tmp := head
    for i, v := range params {
        if i != 0 {
            tmp.Next = &ListNode{Val: v}
            tmp = tmp.Next
        }
    }
    return head, tmp
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

    l1, l2, r1 := makeCircleListNode([]int{4, 1}, []int{5, 0, 1}, []int{8, 4, 5})
    l3, l4, r2 := makeCircleListNode([]int{0, 9, 1}, []int{3}, []int{2, 4})
    l5, l6, r3 := makeCircleListNode([]int{2, 6, 4}, []int{1, 5}, []int{})
    qs := []question{
        {
            p: param{
                one: l1,
                two: l2,
            },
            a: result{
                one: r1,
            },
        },
        {
            p: param{
                one: l3,
                two: l4,
            },
            a: result{
                one: r2,
            },
        },
        {
            p: param{
                one: l5,
                two: l6,
            },
            a: result{
                one: r3,
            },
        },
    }

    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, getIntersectionNode1(p.one, p.two), "输入:%v", p)
        ast.Equal(a.one, getIntersectionNode(p.one, p.two), "输入:%v", p)
    }
}
