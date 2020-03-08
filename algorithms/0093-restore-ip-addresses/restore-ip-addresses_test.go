package restore_ip_addresses

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type param struct {
    one string
}

type result struct {
    one []string
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
                one: "25525511135",
            },
            a: result{
                one: []string{
                    "255.255.11.135",
                    "255.255.111.35",
                },
            },
        },
        {
            p: param{
                one: "0000",
            },
            a: result{
                one: []string{
                    "0.0.0.0",
                },
            },
        },
        {
            p: param{
                one: "010010",
            },
            a: result{
                one: []string{
                    "0.10.0.10",
                    "0.100.1.0",
                },
            },
        },
        {
            p: param{
                one: "172162541",
            },
            a: result{
                one: []string{
                    "17.216.25.41",
                    "17.216.254.1",
                    "172.16.25.41",
                    "172.16.254.1",
                    "172.162.5.41",
                    "172.162.54.1",
                },
            },
        },
        {
            p: param{
                one: "000256",
            },
            a: result{
                one: []string{},
            },
        },
        {
            p: param{
                one: "66212274",
            },
            a: result{
                one: []string{
                    "6.62.122.74",
                    "66.2.122.74",
                    "66.21.22.74",
                    "66.21.227.4",
                    "66.212.2.74",
                    "66.212.27.4",
                },
            },
        },
    }
    for _, q := range qs {
        a, p := q.a, q.p
        ast.Equal(a.one, restoreIpAddresses(p.one), "输入:%v", q)
    }
}
