## 问题
A linked list is given such that each node contains an additional random pointer which could point to any node in the list or null.

Return a deep copy of the list.

The Linked List is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:

- val: an integer representing Node.val
- random_index: the index of the node (range from 0 to n-1) where random pointer points to, or null if it does not point to any node.
 

Example 1:
![img](https://assets.leetcode.com/uploads/2019/12/18/e1.png)
```
Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
```

Example 2:
![img](https://assets.leetcode.com/uploads/2019/12/18/e2.png)
```
Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]
```

Example 3:
![img](https://assets.leetcode.com/uploads/2019/12/18/e3.png)
```
Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]
```

Example 4:
```
Input: head = []
Output: []
Explanation: Given linked list is empty (null pointer), so return null.
```

Constraints:

- -10000 <= Node.val <= 10000
- Node.random is null or pointing to a node in the linked list.
- Number of Nodes will not exceed 1000.

## 分析
给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。

要求返回这个链表的 深拷贝。

感觉没啥，用一个map记录随机指针的值即可，感觉好麻烦，想一个简单的方式
- 新增一个node，然后移动，再拆分

## 最高分
使用map记录
```golang
func copyRandomList(head *Node) *Node {
	cur := head
	dummy := &Node{}
	dCur := dummy
	visited := map[*Node]*Node{}
	// copy list first
	for cur != nil {
		n := &Node{Val: cur.Val}
		visited[cur] = n
		dCur.Next = n
		dCur = dCur.Next
		cur = cur.Next
	}
	// copy random pointers
	cur, dCur = head, dummy
	for cur != nil {
		dCur.Next.Random = visited[cur.Random]
		dCur = dCur.Next
		cur = cur.Next
	}
	return dummy.Next
}
```

## 实现
```golang
func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    var newHead, point *Node

    // insert new node
    for point = head; point != nil; point = newHead.Next {
        newHead = &Node{
            Val:    point.Val,
            Random: point.Random,
        }
        point.Next, newHead.Next = newHead, point.Next
    }

    // assign random node
    for point = head; point != nil; point = point.Next.Next {
        if point.Random != nil {
            point.Next.Random = point.Random.Next
        }
    }

    // split the new Node and original Node
    newHead, point = head.Next, head.Next
    for head != nil {
        if head.Next != nil {
            head.Next = head.Next.Next
        }
        if point.Next != nil {
            point.Next = point.Next.Next
        }
        head, point = head.Next, point.Next
    }
    return newHead
}
```

## 改进
```golang

```

## 反思

## 参考