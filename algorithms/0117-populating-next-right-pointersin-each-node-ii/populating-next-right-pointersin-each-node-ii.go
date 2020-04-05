package populating_next_right_pointersin_each_node_ii

type Node struct {
    Val   int
    Left  *Node
    Right *Node
    Next  *Node
}

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

func connect1(root *Node) *Node {
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
    root.Right, root.Left = connect1(root.Right), connect1(root.Left)
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