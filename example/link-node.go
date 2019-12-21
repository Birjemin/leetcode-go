package main

// 参考: https://segmentfault.com/a/1190000007442806
// 链表实现
import (
	"fmt"
)

const (
	ERROR = -1000000001
)

type Element int64

type LinkNode struct {
	Data Element
	Nest *LinkNode
}

func NewLinkNode() ILinkNode {
	return &LinkNode{0, nil}
}

type ILinkNode interface {
	Add(data Element)
	Delete(index int) Element
	Insert(index int, data Element)
	GetLength() int
	Search(data Element)
	GetData(index int) Element
	Traverse()
}

func (head *LinkNode) Add(data Element) {
	point := head //临时指针
	// 移动到最后一位
	for point.Nest != nil {
		point = point.Nest //移位
	}
	var node LinkNode  //新节点
	point.Nest = &node //赋值
	node.Data = data
}

func (head *LinkNode) Delete(index int) Element {
	//判断index合法性
	if index < 0 || index > head.GetLength() {
		fmt.Println("please check index")
		return ERROR
	} else {
		point := head
		for i := 0; i < index-1; i++ {
			point = point.Nest //移位
		}
		point.Nest = point.Nest.Nest //赋值
		data := point.Nest.Data
		return data
	}
}

func (head *LinkNode) Insert(index int, data Element) {
	//检验index合法性
	if index < 0 || index > head.GetLength() {
		fmt.Println("please check index")
	} else {
		point := head
		for i := 0; i < index-1; i++ {
			point = point.Nest
		}
		var node LinkNode //新节点，赋值
		node.Data = data
		node.Nest = point.Nest
		point.Nest = &node
	}
}

func (head *LinkNode) GetLength() int {
	point := head
	var length int
	for point.Nest != nil {
		length++
		point = point.Nest
	}
	return length
}

func (head *LinkNode) Search(data Element) {
	point := head
	index := 0
	for point.Nest != nil {
		if point.Data == data {
			fmt.Println(data, "exist at", index, "th")
			break
		} else {
			index++
			point = point.Nest
			if index > head.GetLength()-1 {
				fmt.Println(data, "not exist at")
				break
			}
			continue
		}
	}
}

func (head *LinkNode) GetData(index int) Element {
	point := head
	if index < 0 || index > head.GetLength() {
		fmt.Println("please check index")
		return ERROR
	} else {
		for i := 0; i < index; i++ {
			point = point.Nest
		}
		return point.Data
	}
}

func (head *LinkNode) Traverse() {
	point := head.Nest
	for point.Nest != nil {
		fmt.Println(point.Data)
		point = point.Nest
	}
	fmt.Println("Traverse OK!")
}

func main() {
	var head = NewLinkNode()
	// 添加十个数字
	for i := 0; i < 10; i++ {
		head.Add(Element(i + 1 + i*100))
	}

	head.Traverse()

	head.Delete(3)
	head.Traverse()

	head.Search(809)
	head.Traverse()

	head.Insert(4, 10010)
	head.Traverse()
}
