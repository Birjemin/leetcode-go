## 问题
Write a program to find the node at which the intersection of two singly linked lists begins.

For example, the following two linked lists:

![img](https://assets.leetcode.com/uploads/2018/12/13/160_statement.png)

begin to intersect at node c1.

Example 1:
![img](https://assets.leetcode.com/uploads/2018/12/13/160_example_1.png)
```
Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
Output: Reference of the node with value = 8
Input Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,0,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.
```

Example 2:
![img](https://assets.leetcode.com/uploads/2018/12/13/160_example_2.png)
```
Input: intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
Output: Reference of the node with value = 2
Input Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [0,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.
```

Example 3:
![img](https://assets.leetcode.com/uploads/2018/12/13/160_example_3.png)
```
Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
Output: null
Input Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
Explanation: The two lists do not intersect, so return null.
```

Notes:

- If the two linked lists have no intersection at all, return null.
- The linked lists must retain their original structure after the function returns.
- You may assume there are no cycles anywhere in the entire linked structure.
- Your code should preferably run in O(n) time and use only O(1) memory.

## 分析
编写一个程序，找到两个单链表相交的起始节点。
- 最简单的就是两个遍历
- 就是使用空间换取时间，保存遍历headA结果,然后遍历headB
- 比较高端，双指针法，交换两个指针，headA遍历完了则连接headB继续遍历，反之类似，然后相遇则表示相交。

## 最高分
```golang
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    //boundary check
    if headA == nil || headB == nil {
        return nil
    }

    a, b := headA, headB

    //if a & b have different len, then we will stop the loop after second iteration
    for a != b {
        //for the end of first iteration, we just reset the pointer to the head of another linkedlist
        if a == nil {
            a = headB
        } else {
            a = a.Next
        }

        if b == nil {
            b = headA
        } else {
            b = b.Next
        }
    }
    return a
}
```

## 实现
最直接的实现方式
```golang
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    for ; headA != nil; headA = headA.Next {
        for tmp := headB; tmp != nil; tmp = tmp.Next {
            if headA == tmp {
                return headA
            }
        }
    }
    return nil
}
```

## 改进
太慢，更改方式
```golang
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    var end *ListNode
    tmpA, tmpB := headA, headB

    for tmpA != tmpB {
        if tmpA.Next == nil {
            if end == nil {
                end = tmpA
            } else if end != tmpA {
                return nil
            }
            tmpA = headB
        } else {
            tmpA = tmpA.Next
        }
        if tmpB.Next == nil {

            if end == nil {
                end = tmpB
            } else if end != tmpB {
                return nil
            }
            tmpB = headA
        } else {
            tmpB = tmpB.Next
        }
    }
    return tmpA
}
```

## 改进
还慢，更改方式，直接判断当前节点
```golang
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    tmpA, tmpB := headA, headB

    for tmpA != tmpB {
        if tmpA == nil {
            tmpA = headB
        } else {
            tmpA = tmpA.Next
        }
        if tmpB == nil {
            tmpB = headA
        } else {
            tmpB = tmpB.Next
        }
    }
    return tmpA
}
```

## 反思

## 参考