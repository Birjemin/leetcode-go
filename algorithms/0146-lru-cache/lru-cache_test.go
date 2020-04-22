package lru_cache

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test(t *testing.T) {
    ast := assert.New(t)

    cache := Constructor(1)
    cache.Put(2, 1)

    ast.Equal(1, cache.Get(2), "输入:%v")
    //
    cache.Put(3, 2) // 该操作会使得密钥 1 作废
    //
    ast.Equal(-1, cache.Get(2), "输入:%v")
    //
    ast.Equal(2, cache.Get(3), "输入:%v")
}
