package binary_search_tree_iterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
