## 问题
Write a function that takes an unsigned integer and return the number of '1' bits it has (also known as the Hamming weight).

Example 1:
```
Input: 00000000000000000000000000001011
Output: 3
Explanation: The input binary string 00000000000000000000000000001011 has a total of three '1' bits.
```

Example 2:
```
Input: 00000000000000000000000010000000
Output: 1
Explanation: The input binary string 00000000000000000000000010000000 has a total of one '1' bit.
```

Example 3:
```
Input: 11111111111111111111111111111101
Output: 31
Explanation: The input binary string 11111111111111111111111111111101 has a total of thirty one '1' bits.
```

Note:

- Note that in some languages such as Java, there is no unsigned integer type. In this case, the input will be given as signed integer type and should not affect your implementation, as the internal binary representation of the integer is the same whether it is signed or unsigned.
- In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 3 above the input represents the signed integer -3.

Follow up:

If this function is called many times, how would you optimize it?

## 分析
编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）


## 最高分
```golang
func hammingWeight(num uint32) int {
    count := 0
    for num != 0 {
        count++
        num &= num - 1
    }
    return count
}
```

## 实现
最简单的方式
```golang
func hammingWeight(num uint32) int {
    var count int
    for i := 0; i < 32; i++ {
        if num&1 == 1 {
            count++
        }
        num >>= 1
    }
    return count
}
```

## 改进
```golang
func hammingWeight(num uint32) int {
    var count int
    for ; num >= 1; num >>= 1 {
        if num&1 == 1 {
            count++
        }
    }
    return count
}
```

## 反思

## 参考