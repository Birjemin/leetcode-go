## 问题
Given a binary tree
```
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
```

Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.

 

Follow up:

- You may only use constant extra space.
- Recursive approach is fine, you may assume implicit stack space does not count as extra space for this problem.
 

Example 1:

![1](https://assets.leetcode.com/uploads/2019/02/15/117_sample.png)
```
Input: root = [1,2,3,4,5,null,7]
Output: [1,#,2,3,#,4,5,7,#]
Explanation: Given the above binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.
```

Constraints:

- The number of nodes in the given tree is less than 6000.
- `-100 <= node.val <= 100`

## 分析
和116题一样

## 最高分
```golang

```

## 实现
bfs算法
```golang
func connect(root *Node) *Node {
    if root == nil {
        return nil
    }
    bfs([]*Node{root})
    return root
}

func bfs(tmp []*Node) {
    if tmp == nil {
        return
    }
    var ret []*Node
    var p *Node

    for _, v := range tmp {
        if p != nil {
            p.Next = v
        }
        p = v
        if v.Left != nil {
            ret = append(ret, v.Left)
        }
        if v.Right != nil {
            ret = append(ret, v.Right)
        }
    }
    bfs(ret)
}
```

## 改进
用法挺精巧的~可以学习一下~~
```golang
func connect(root *Node) *Node {
    if root == nil || (root.Left == nil && root.Right == nil) {
        return root
    }
    if root.Left != nil && root.Right != nil {
        root.Left.Next, root.Right.Next = root.Right, getNextChildNotNull(root)
    } else if root.Right == nil {
        root.Left.Next = getNextChildNotNull(root)
    } else {
        root.Right.Next = getNextChildNotNull(root)
    }
    root.Right, root.Left = connect(root.Right), connect(root.Left)
    return root
}

func getNextChildNotNull(n *Node) *Node {
    for ; n.Next != nil; n = n.Next {
        if n.Next.Left != nil {
            return n.Next.Left
        }
        if n.Next.Right != nil {
            return n.Next.Right
        }
    }
    return nil
}

```

## 反思

## 参考