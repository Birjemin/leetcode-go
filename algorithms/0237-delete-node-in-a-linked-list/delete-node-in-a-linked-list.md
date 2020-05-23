## 问题
Write a function to delete a node (except the tail) in a singly linked list, given only access to that node.

Given linked list -- head = [4,5,1,9], which looks like following:

![img](https://assets.leetcode.com/uploads/2018/12/28/237_example.png)

Example 1:
```
Input: head = [4,5,1,9], node = 5
Output: [4,1,9]
Explanation: You are given the second node with value 5, the linked list should become 4 -> 1 -> 9 after calling your function.
```

Example 2:
```
Input: head = [4,5,1,9], node = 1
Output: [4,5,9]
Explanation: You are given the third node with value 1, the linked list should become 4 -> 5 -> 9 after calling your function.
```

Note:

- The linked list will have at least two elements.
- All of the nodes' values will be unique.
- The given node will not be the tail and it will always be a valid node of the linked list.
- Do not return anything from your function.

## 分析
请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。

乍一看，还真的没思路。。。

## 最高分
这个就很流弊了。
```golang
func deleteNode(node *ListNode) {
    *node = *node.Next
}
```

## 实现
后面覆盖前面的值
```golang
func deleteNode(node *ListNode) {
	for node != nil {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			node.Next = nil
		}
		node = node.Next
	}
}
```

## 改进
只要把后面那一位的值复制到前面去，然后连接第三个值就好了
```golang
func deleteNode(node *ListNode) {
	node.Val, node.Next = node.Next.Val, node.Next.Next
}
```

## 反思

## 参考