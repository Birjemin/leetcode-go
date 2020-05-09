## 问题
Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.

Calling next() will return the next smallest number in the BST.

Example:

![img](https://assets.leetcode.com/uploads/2018/12/25/bst-tree.png)

```
BSTIterator iterator = new BSTIterator(root);
iterator.next();    // return 3
iterator.next();    // return 7
iterator.hasNext(); // return true
iterator.next();    // return 9
iterator.hasNext(); // return true
iterator.next();    // return 15
iterator.hasNext(); // return true
iterator.next();    // return 20
iterator.hasNext(); // return false
```

Note:

- next() and hasNext() should run in average O(1) time and uses O(h) memory, where h is the height of the tree.
- You may assume that next() call will always be valid, that is, there will be at least a next smallest number in the BST when next() is called.

## 分析
实现二叉搜索树迭代器
- 最简单的方式就是中序遍历保存结果(不满足要求)

## 最高分
也不符合要求
```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    sorted []int
}


func Constructor(root *TreeNode) BSTIterator {
    
    var iter BSTIterator
    iter.inOrder(root)
    
    return iter
}

func (this *BSTIterator) inOrder(cur *TreeNode) {
    
    if cur == nil {
        return
    }
    
    this.inOrder(cur.Left)
    this.sorted = append(this.sorted, cur.Val)
    this.inOrder(cur.Right)
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    
    n := this.sorted[0]
    this.sorted = this.sorted[1:]
    return n
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    
    return len(this.sorted) != 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
```

## 实现
可以使用中序遍历来获取数据（不符合要求）
```golang
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	tmp []int
	length int
}

// inorder
func cal(t *TreeNode, ret *[]int) {
	if t == nil {
		return
	}
	// left child node
	cal(t.Left, ret)
	*ret = append(*ret, t.Val)
	// right child node
	cal(t.Right, ret)
}

func Constructor(root *TreeNode) BSTIterator {
	obj := BSTIterator{tmp: []int{},length:0}
	cal(root, &obj.tmp)
	obj.length = len(obj.tmp)
	return obj
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if this.length <= 0 {
		return 0
	}
	val := this.tmp[0]
	this.tmp = this.tmp[1:]
	this.length--
	return val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.length <= 0 {
		return false
	}
	return true
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

```

## 改进
这个方法真刘辟~~~~
主要是考虑到遍历的关系，使用栈，先将left节点保存，然后取的时候将right节点推入到栈中，刘辟刘辟~~
```
type BSTIterator struct {
	stack  []*TreeNode
	length int
}

func Constructor(root *TreeNode) BSTIterator {
	obj := BSTIterator{}
	obj.push(root)
	return obj
}

func (this *BSTIterator) push(root *TreeNode) {
	for ; root != nil; root = root.Left {
		this.length++
		this.stack = append(this.stack, root)
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	top := this.stack[this.length-1]
	this.stack = this.stack[:this.length-1]
	this.length--
	this.push(top.Right)
	return top.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.length != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

```

## 反思

## 参考