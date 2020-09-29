package main

import (
	"fmt"
	"github.com/mihir1003/DSSE/BinaryTrees"
	"github.com/mihir1003/DSSE/Lists"
)

func main() {

	myList := &Lists.LinkedList{Val: 5, Next: nil}

	myList.add(5)

	myList.add(9)

	myList.add(6)

	myList.add(7)

	myList.delete(0)

	myListSorted := myList.sort()
	fmt.Println((*myList).extractMin())
	myList.display()
	myListSorted.display()

	myTree := BinaryTrees.ItemBinarySearchTree{}
	myTree.Insert(4, 1)
	myTree.Insert(9, 1)
	myTree.Insert(10, 6)
	myTree.Insert(5, 9)
	fmt.Println(*myTree.Max())

}
