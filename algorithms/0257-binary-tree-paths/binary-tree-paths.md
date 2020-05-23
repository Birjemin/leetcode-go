## 问题
Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.

Example:
```
Input:

   1
 /   \
2     3
 \
  5

Output: ["1->2->5", "1->3"]

Explanation: All root-to-leaf paths are: 1->2->5, 1->3
```

## 分析
给定一个二叉树，返回所有从根节点到叶子节点的路径。

dfs、bfs算法都可以

## 最高分
```golang
var arrow string = "->"

func binaryTreePaths(root *TreeNode) []string {
    ans := make([]string, 0)
    
    paths(root, "", &ans)
    
    return ans
}

func paths(root *TreeNode, prefix string, ans *[]string) {
    if root == nil {
        return
    }
    
    if len(prefix) == 0 {
        prefix += strconv.Itoa(root.Val)
    } else {
        prefix += "->" + strconv.Itoa(root.Val)
    }
    
    if root.Left == nil && root.Right == nil {
        *ans = append(*ans, prefix)
        return
    }
    
    paths(root.Left, prefix, ans)
    paths(root.Right, prefix, ans)
}
```

## 实现
```golang
func binaryTreePaths(root *TreeNode) []string {
	var ret []string
	dfs(root, "", &ret)
	return ret
}

func dfs(t *TreeNode, str string, ret *[]string) {
	if t == nil {
		return
	}
	str += strconv.Itoa(t.Val)
	if t.Left == nil && t.Right == nil {
		*ret = append(*ret, str)
		return
	}
	str += "->"
	dfs(t.Left, str, ret)
	dfs(t.Right, str, ret)
}

```

## 改进
```golang

```

## 反思

## 参考