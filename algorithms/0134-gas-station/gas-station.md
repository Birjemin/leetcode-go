## 问题
There are N gas stations along a circular route, where the amount of gas at station i is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from station i to its next station (i+1). You begin the journey with an empty tank at one of the gas stations.

Return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1.

Note:

* If there exists a solution, it is guaranteed to be unique.
* Both input arrays are non-empty and have the same length.
* Each element in the input arrays is a non-negative integer.

Example 1:
```
Input: 
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]

Output: 3

Explanation:
Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 4. Your tank = 4 - 1 + 5 = 8
Travel to station 0. Your tank = 8 - 2 + 1 = 7
Travel to station 1. Your tank = 7 - 3 + 2 = 6
Travel to station 2. Your tank = 6 - 4 + 3 = 5
Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
Therefore, return 3 as the starting index.
```

Example 2:
```
Input: 
gas  = [2,3,4]
cost = [3,4,3]

Output: -1

Explanation:
You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 0. Your tank = 4 - 3 + 2 = 3
Travel to station 1. Your tank = 3 - 3 + 3 = 3
You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
Therefore, you can't travel around the circuit once no matter where you start.
```

## 分析
在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。

- 不用考虑多解问题

## 最高分
```golang
func canCompleteCircuit(gas []int, cost []int) int {
    start := 0
    total := 0
    tank := 0
    for i := 0; i < len(gas); i++ {
        tank += gas[i] - cost[i]
        if tank < 0 {
            start = i + 1
            total += tank
            tank = 0
        }
    }
    if total + tank < 0 {
        return -1
    }
    return start
}
```

## 实现
最简单的方式
```golang
func canCompleteCircuit(gas []int, cost []int) int {

    length := len(gas)

    // get subtraction
    for i := range gas {
        gas[i] -= cost[i]
    }

    for i, v := range gas {
        // next value
        if v < 0 {
            continue
        }

        for j := 1; j <= length; j++ {
            // got it
            if j == length {
                return i
            }

            p := j + i
            if p >= length {
                p -= length
            }

            v += gas[p]
            if v < 0 {
                break
            }
        }
    }

    return -1
}

```

## 改进
巧妙的方式，需要想一想
```golang
func canCompleteCircuit(gas []int, cost []int) int {
    total, curr, start := 0, 0, 0
    for i, v := range gas {
        // v - cost
        total += v - cost[i]
        curr += v - cost[i]

        if curr < 0 {
            start, curr = i+1, 0
        }
    }

    // got if
    if total >= 0 {
        return start
    }

    return -1
}
```

## 反思

## 参考