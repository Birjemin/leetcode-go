## 问题
Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example:

Given the sorted linked list: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

```
      0
     / \
   -3   9
   /   /
 -10  5
```

## 分析
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

二叉搜索树中序遍历是一个有序的数组

- 转成数组，然后递归处理
- 快慢指针，递归处理
- 模拟中序方式（其实和第二种类似，无非是一个计算长度得到中间值，一个是移动指针得到中间值）

## 最高分
```golang

```

## 实现
转成数组
```golang
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    temp := toArr(head)
    return cal(temp)
}

func cal(arr []int) *TreeNode {
    length := len(arr)
    if length == 0 {
        return nil
    }
    mid := length / 2
    return &TreeNode{
        Val:   arr[mid],
        Left:  cal(arr[:mid]),
        Right: cal(arr[mid+1:]),
    }
}

func toArr(head *ListNode) (ret []int) {
    for ; head != nil; head = head.Next {
        ret = append(ret, head.Val)
    }
    return
}
```

## 改进
快慢指针
```golang
func sortedListToBST(head *ListNode) *TreeNode {
    idx := findIndex(head)
    if idx == nil {
        return nil
    }
    if idx.Next == nil {
        return &TreeNode{Val: idx.Val}
    }

    val, n := idx.Next.Val, idx.Next.Next
    idx.Next = nil

    return &TreeNode{
        Val:   val,
        Left:  sortedListToBST(head),
        Right: sortedListToBST(n)}
}

func findIndex(slow *ListNode) *ListNode {
    if slow == nil {
        return nil
    }
    if slow.Next == nil {
        return slow
    }
    fast := slow.Next
    for {
        fast = fast.Next
        if fast == nil || fast.Next == nil {
            return slow
        }
        fast, slow = fast.Next, slow.Next
    }
}
```

## 反思

## 参考