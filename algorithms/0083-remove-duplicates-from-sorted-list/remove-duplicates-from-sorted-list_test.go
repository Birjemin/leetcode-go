package remove_duplicates_from_sorted_list

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
               one: makeListNode([]int{1, 1, 2}),
           },
           a: result{
               one: makeListNode([]int{1, 2}),
           },
        },
        {
           p: param{
               one: makeListNode([]int{1, 1, 2, 3, 3}),
           },
           a: result{
               one: makeListNode([]int{1, 2, 3}),
           },
        },
        {
          p: param{
              one: makeListNode([]int{0, 0, 0, 0, 3}),
          },
          a: result{
              one: makeListNode([]int{0, 3}),
          },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, deleteDuplicates2(p.one), "输入:%v", q)
        ast.Equal(a.one, deleteDuplicates1(p.one), "输入:%v", q)
        ast.Equal(a.one, deleteDuplicates(p.one), "输入:%v", q)
    }
}