package best_time_to_buy_and_sell_stock_ii

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one []int
}

type result struct {
    one int
}

type question struct {
    p param
    a result
}

func Test(t *testing.T) {
    ast := assert.New(t)
    qs := []question{
        {
            p: param{
                one: []int{7, 1, 5, 3, 6, 4},
            },
            a: result{
                one: 7,
            },
        },
        {
            p: param{
                one: []int{1, 2, 3, 4, 5},
            },
            a: result{
                one: 4,
            },
        },
        {
            p: param{
                one: []int{7, 6, 4, 3, 1},
            },
            a: result{
                one: 0,
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, maxProfit(p.one), "输入:%v", q)
        ast.Equal(a.one, maxProfit1(p.one), "输入:%v", q)
    }
}
