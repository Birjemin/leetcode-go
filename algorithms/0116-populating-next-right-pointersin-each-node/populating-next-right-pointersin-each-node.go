package populating_next_right_pointersin_each_node

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
            ret = append(ret, v.Left, v.Right)
        }
    }
    bfs(ret)
}

func connect1(root *Node) *Node {
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
