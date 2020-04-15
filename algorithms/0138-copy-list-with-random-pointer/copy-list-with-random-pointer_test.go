package copy_list_with_random_pointer

import (
    "fmt"
    "github.com/spf13/cast"
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one *Node
}

type result struct {
    one *Node
}

type question struct {
    p param
    a result
}

func makeListNode(params [][]int) *Node {
    length := len(params)
    if length == 0 {
        return nil
    }
    randArr := make([]int, length)
    nodeArr := make([]*Node, length)

    res := &Node{Val: params[0][0]}
    randArr[0] = params[0][1]
    nodeArr[0] = res
    temp := res

    for i, v := range params {
        if i != 0 {
            temp.Next = &Node{Val: v[0]}
            randArr[i] = v[1]
            temp = temp.Next
            nodeArr[i] = temp
        }
    }

    for i, v := range randArr {
        if v == -1 {
            nodeArr[i].Random = nil
        } else {
            nodeArr[i].Random = nodeArr[v]
        }
    }
    return res
}

func showListNode(res *Node) {
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
                one: makeListNode([][]int{
                    {7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0},
                }),
            },
            a: result{
                one: makeListNode([][]int{
                    {7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0},
                }),
            },
        },
        {
            p: param{
                one: makeListNode([][]int{
                    {1, 1}, {2, 1},
                }),
            },
            a: result{
                one: makeListNode([][]int{
                    {1, 1}, {2, 1},
                }),
            },
        },
        {
            p: param{
                one: makeListNode([][]int{
                    {3, -1}, {3, 0}, {3, -1},
                }),
            },
            a: result{
                one: makeListNode([][]int{
                    {3, -1}, {3, 0}, {3, -1},
                }),
            },
        },
        {
            p: param{
                one: makeListNode([][]int{}),
            },
            a: result{
                one: makeListNode([][]int{}),
            },
        },
    }

    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, copyRandomList(p.one), "输入:%v", p)
    }
}
