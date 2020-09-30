package main

import (
	"fmt"
	BinaryTrees "github.com/mihir1003/DSSE/dataStructure/BinaryTrees"
	Lists "github.com/mihir1003/DSSE/dataStructure/Lists"
)

func main() {

	myList := &Lists.LinkedList{Val: 9, Next: nil}
	myList.Add(9)
	myList.Display()
	myTree := BinaryTrees.ItemBinarySearchTree{}
	myTree.Insert(4, 1)
	myTree.Insert(9, 1)
	myTree.Insert(10, 6)
	myTree.Insert(5, 9)
	fmt.Println(*myTree.Max())

}
