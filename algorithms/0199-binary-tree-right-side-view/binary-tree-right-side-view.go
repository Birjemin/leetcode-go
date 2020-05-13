package binary_tree_right_side_view

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func rightSideView1(root *TreeNode) []int {
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

func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	tmp := list.New()
	tmp.PushBack(root)

	var ret []int
	var val int

	for {
		length := tmp.Len()
		if length == 0 {
			break
		}

		for ;length > 0;length-- {
			e := tmp.Back()
			tmp.Remove(e)
			n := e.Value.(*TreeNode)
			if n.Left != nil {
				tmp.PushFront(n.Left)
			}
			if n.Right != nil {
				tmp.PushFront(n.Right)
			}
			val = n.Val
		}
		ret = append(ret, val)
	}

	return ret
}
