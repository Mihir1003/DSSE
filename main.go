package main

import (
	"fmt"
	"github.com/mihir1003/DSSE/BinaryTrees"
	"github.com/mihir1003/DSSE/Lists"
)

func main() {

	myList := &Lists.LinkedList{Val: 5, Next: nil}

	myList.Add(5)

	myList.Add(9)

	myList.Add(6)

	myList.Add(7)

	myList.Delete(0)

	myListSorted := myList.Sort()
	fmt.Println((*myList).ExtractMin())
	myList.Display()
	myListSorted.Display()

	myTree := BinaryTrees.ItemBinarySearchTree{}
	myTree.Insert(4, 1)
	myTree.Insert(9, 1)
	myTree.Insert(10, 6)
	myTree.Insert(5, 9)
	fmt.Println(*myTree.Max())

}
