package copy_list_with_random_pointer

type Node struct {
    Val    int
    Next   *Node
    Random *Node
}

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
