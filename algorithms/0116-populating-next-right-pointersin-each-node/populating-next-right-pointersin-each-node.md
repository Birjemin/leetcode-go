## 问题
You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:
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

```
Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]
Explanation: Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.
 ```

Constraints:

- The number of nodes in the given tree is less than 4096.
- `-1000 <= node.val <= 1000`

## 分析
给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
```
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
```
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

![示意图](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/02/15/116_sample.png)

bfs不可使用，因为限制了常数级空间，可以使用变种

测试用例太难写了

## 最高分
```golang
func connect(root *Node) *Node {
    if root == nil {
        return nil
    }
    for l := root; l.Left != nil; l = l.Left {
        for r := l; r.Next != nil; r = r.Next {
            r.Left.Next = r.Right
            if r.Right != nil {
                r.Right.Next = r.Next.Left
            }
        }
    }
    return root
}
```

## 实现
bfs 需要额外空间，不是最优解
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
            ret = append(ret, v.Left, v.Right)
        }
    }
    bfs(ret)
}
```

## 改进
参考最高分，进行解答，这个解答有点意思，也挺有趣的
```golang
func connect(root *Node) *Node {
    if root == nil {
        return nil
    }
    for l := root; l.Left != nil; l = l.Left {
        for r := l; r.Next != nil; r = r.Next {
            r.Left.Next = r.Right
            if r.Right != nil {
                r.Right.Next = r.Next.Left
            }
        }
    }
    return root
}
```

## 反思

## 参考