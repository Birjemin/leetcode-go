## 问题

![img](https://upload.wikimedia.org/wikipedia/commons/0/0f/Insertion-sort-example-300px.gif)

A graphical example of insertion sort. The partial sorted list (black) initially contains only the first element in the list.
With each iteration one element (red) is removed from the input data and inserted in-place into the sorted list
 

Algorithm of Insertion Sort:

Insertion sort iterates, consuming one input element each repetition, and growing a sorted output list.
At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list, and inserts it there.
It repeats until no input elements remain.

Example 1:
```
Input: 4->2->1->3
Output: 1->2->3->4
```

Example 2:
```
Input: -1->5->3->4->0
Output: -1->0->3->4->5
```

## 分析
对链表进行插入排序，注意移动就好

## 最高分

```golang
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{Val: 0, Next: nil} // 这里初始化不要直接指向 head，为了下面循环可以统一处理
	cur := head
	pre := newHead
	next := &ListNode{Val: 0, Next: nil}
	for cur != nil {
		next = cur.Next
		for pre.Next != nil && pre.Next.Val < cur.Val {
			pre = pre.Next
		}
		cur.Next = pre.Next
		pre.Next = cur
		pre = newHead // 归位，重头开始
		cur = next
	}
	return newHead.Next
}
```

## 实现
先解决插入一个的问题，然后再解决多个的问题（效率不高）
```golang
func insertionSortList(head *ListNode) *ListNode {

    if head == nil {
        return nil
    }

    tmp := head.Next

    for head.Next = nil; tmp != nil; {
        t := tmp.Next
        tmp.Next = nil
        head = insertNode(head, tmp)
        tmp = t
    }
    return head
}

// insert one node
func insertNode(start, target *ListNode) *ListNode {
    // first
    var front *ListNode
    head := start
    for start != nil {
        if start.Val > target.Val {
            target.Next = start
            break
        }
        front, start = start, start.Next
    }
    if front == nil {
        head = target
    } else {
        front.Next = target
    }
    return head
}

```

## 改进
两次循环解决这个问题(超时)
```golang
func insertionSortList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    // add new head
    newHead := &ListNode{Val: 0, Next: head}

    var tmp *ListNode
    // current node
    for curr := head; curr.Next != nil; curr = tmp {
        // storage the next of current node
        tmp = curr.Next

        for pre := newHead; pre != nil && pre != curr; pre = pre.Next {
            if pre.Next != nil && pre.Next.Val > curr.Next.Val {
                // insert the node
                curr.Next.Next, pre.Next, curr.Next = pre.Next, curr.Next, curr.Next.Next
                break
            }
        }
    }
    return newHead.Next
}
```

## 改进
使用一个新的指针，有序集合，然后遍历源数据并且插入相应的位置（方法很奇妙~~）
```golang
func insertionSortList1(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    var tmp *ListNode
    tummy := &ListNode{}
    for pre := tummy; head != nil; head = tmp {

        // store the the next node of head(head=tmp)
        // reset the first node
        pre, tmp = tummy, head.Next

        // find the position should be exchanged
        for pre.Next != nil && pre.Next.Val < head.Val {
            pre = pre.Next
        }

        // insert the node
        head.Next, pre.Next = pre.Next, head
    }
    return tummy.Next
}
```

## 反思

## 参考