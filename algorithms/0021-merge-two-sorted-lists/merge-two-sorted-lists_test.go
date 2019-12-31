package merge_two_sorted_lists

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
                one: makeListNode([]int{1, 2, 4}),
                two: makeListNode([]int{1, 3, 4}),
            },
            a: result{
                one: makeListNode([]int{1, 1, 2, 3, 4, 4}),
            },
        },
        {
           p: param{
               one: makeListNode([]int{}),
               two: makeListNode([]int{}),
           },
           a: result{
               one: makeListNode([]int{}),
           },
        },
        {
           p: param{
               one: makeListNode([]int{}),
               two: makeListNode([]int{1, 3, 4}),
           },
           a: result{
               one: makeListNode([]int{1, 3, 4}),
           },
        },
        {
           p: param{
               one: makeListNode([]int{1, 2}),
               two: makeListNode([]int{9, 9, 9}),
           },
           a: result{
               one: makeListNode([]int{1, 2, 9, 9, 9}),
           },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, mergeTwoLists(p.one, p.two), "输入:%v", q)
        //ast.Equal(a.one, mergeTwoLists1(p.one, p.two), "输入:%v", q)
    }
}
