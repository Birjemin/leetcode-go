## 问题
Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

Example:
```
Input: [1,2,3,null,5,null,4]
Output: [1, 3, 4]
Explanation:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
```

## 分析
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
最简单的方式就是使用bfs算法，102题和107题类似

## 最高分
```golang
func rightSideView(root *TreeNode) []int {
    if root == nil { return []int{} }
    
    queue := list.New()
    swap := list.New()
    
    ans := []int{}
    
    queue.PushBack(root)
    
    for queue.Len() > 0 {
        // deque
        e := queue.Front()
        queue.Remove(e)
        
        node := e.Value.(*TreeNode)
        if node.Left != nil {
            swap.PushBack(node.Left)
        }
        
        if node.Right != nil {
            swap.PushBack(node.Right)
        }
        
        if queue.Len() == 0 {
            ans = append(ans, node.Val)
            queue, swap = swap, queue
        }
    }
    
    return ans
}
```

## 实现
bfs算法(可以考虑使用内置的container/list实现)
```golang
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	tmp := []*TreeNode{root}
	var ret []int
	var key int

	for {
		var t []*TreeNode
		for _, n := range tmp {
			if n.Left != nil {
				t = append(t, n.Left)
			}
			if n.Right != nil {
				t = append(t, n.Right)
			}
			key = n.Val
		}
		ret = append(ret, key)
		if len(t) == 0 {
			break
		}
		tmp = t
	}
	return ret
}
```

## 改进
尝试使用dfs算法
```golang
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var ret []int
	dfs(root, 0, &ret)
	return ret
}

func dfs(root *TreeNode, length int, ret *[]int) {
	if root == nil {
		return
	}
	if len(*ret) <= length {
		*ret = append(*ret, root.Val)
	}
	dfs(root.Right, length+1, ret)
	dfs(root.Left, length+1, ret)
}
```

## 改进
bfs算法使用bfs算法，使用内置的container/list实现

## 反思

## 参考