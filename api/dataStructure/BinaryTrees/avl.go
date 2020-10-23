package BinaryTrees

import (
	"fmt"
)

/*
** an avl tree api:
//==>> operator tree api:
//insert an x to AvlTree, invoker format: t = tree.Insert(t, x)
func Insert(tree *AvlNode, x int) *AvlNode
//del an x in AvlTree, invoker format: t = tree.delete(t, x)
func delete(tree *AvlNode, x int) *AvlNode
//make a tree to be empty
func MakeEmpty( tree *AvlNode) *AvlNode
//==>> print tree api:
//draw a tree
func DrawTree1(tree *AvlNode)
//leve print a tree
func LevePrintTree(tree *AvlNode)
//middle print a tree
func MidPrintTree(tree *AvlNode)
//prev print a tree
func PrevPrintTree(tree *AvlNode) {
//return the height of the node
func (node *AvlNode) nodeHeight() int
//==>> find   node  api:
//find Element x position in the avl tree
func (tree *AvlNode) Find(x int) *AvlNode
//find min elem in the tree
func (tree *AvlNode) FinMin() *AvlNode
//find max elem in the tree
func (tree *AvlNode) FinMax() *AvlNode
*/

type AvlNode struct {
	Elem   int
	Left   *AvlNode
	Right  *AvlNode
	Height int
}

func If(condition bool, truVal, falseVal interface{}) interface{} {
	if condition {
		return truVal
	}
	return falseVal
}

//make a tree to be empty
func makeEmpty(tree *AvlNode) *AvlNode {
	//free left child
	if tree.Left != nil {
		tree.Left = makeEmpty(tree.Left)
	}
	//free right child
	if tree.Right != nil {
		tree.Right = makeEmpty(tree.Right)
	}
	//a leaf, just return nil, a gc will free the memory
	return nil
}

func printSpace(spaceNum int) {
	for i := 0; i < spaceNum; i++ {
		fmt.Printf("  ")
	}
}

////draw a tree
//func DrawTree1(tree *AvlNode) {
//	var s [] *AvlNode
//	var nextLine  *AvlNode
//	var curH int
//	newLine := true
//
//	for tree != nil {
//		curH = tree.nodeHeight()
//		if nextLine == tree {
//			fmt.Println()
//			//maxLineNum = int(math.Pow(2, float64(curL)))
//			newLine = true
//			nextLine = nil
//		}
//
//		printSpace( int(math.Pow(2, float64(curH-1)) ) )
//
//		if !newLine {
//			if curH == 0 || curH == 1  {
//				printSpace(1)
//			} else {
//				printSpace( int(math.Pow(2, float64(curH-1))) )
//			}
//		} else {
//			newLine = false
//			if curH == 0 {
//				printSpace(1)
//			}
//		}
//
//		fmt.Printf("%v ", tree.Elem)
//
//		if tree.Left != nil {
//			if nextLine == nil {
//				nextLine = tree.Left
//			}
//			s = append(s, tree.Left)
//		}
//		if tree.Right != nil {
//			if nextLine == nil {
//				nextLine = tree.Right
//			}
//			s = append(s, tree.Right)
//		}
//
//		if len(s) > 0 {
//			tree = s[0]
//			s = s[1:]
//		} else {
//			tree = nil
//		}
//	}
//
//	fmt.Println()
//}
//
////leve print a tree
//func LevePrintTree(tree *AvlNode) {
//	var s [] *AvlNode
//	for tree != nil {
//		fmt.Printf("%v ", tree.Elem)
//
//		if tree.Left != nil {
//			s = append(s, tree.Left)
//		}
//		if tree.Right != nil {
//			s = append(s, tree.Right)
//		}
//
//		if len(s) > 0 {
//			tree = s[0]
//			s = s[1:]
//		} else {
//			tree = nil
//		}
//	}
//
//	fmt.Println()
//
//}
//
////middle print a tree
//func MidPrintTree(tree *AvlNode) {
//	if (tree != nil ) {
//		MidPrintTree(tree.Left)
//		fmt.Printf(" %v ", tree.Elem)
//		MidPrintTree(tree.Right)
//	}
//}
////prev print a tree
//func PrevPrintTree(tree *AvlNode) {
//	if (tree != nil ) {
//		fmt.Printf(" %v ", tree.Elem)
//		PrevPrintTree(tree.Left)
//		PrevPrintTree(tree.Right)
//	}
//}

//find Element x position in the avl tree
func (tree *AvlNode) find(x int) *AvlNode {
	if tree != nil {
		if x < tree.Elem {
			//find left
			return tree.Left.find(x)
		} else if x > tree.Elem {
			//find right
			return tree.Right.find(x)
		} else if x == tree.Elem {
			//we find that
			return tree
		}
	}

	//can not found or empty tree
	return nil
}

//find min elem in the tree
func (tree *AvlNode) ExtractMin() int {
	if tree != nil {
		if tree.Left != nil {
			return tree.Left.ExtractMin()
		} else {
			return tree.Elem
		}

	} else {
		return 0
	}
}

//find max elem in the tree
func (tree *AvlNode) ExtractMax() int {
	if tree != nil {
		if tree.Right != nil {
			return tree.Right.ExtractMax()
		} else {
			return tree.Elem
		}

	} else {
		return 0
	}
}

//return the height of the node
func (node *AvlNode) nodeHeight() int {
	if node == nil {
		return -1
	} else {
		return node.Height
	}
}

func (tree *AvlNode) Add(x int) {
	*tree = *add(tree, x)
}

//insert an x to AvlTree, invoker format: t = tree.Insert(t, x)
func add(tree *AvlNode, x int) *AvlNode {
	if tree == nil {
		//tree = new(AvlNode)
		tree = &AvlNode{
			Elem:   x,
			Height: 0,
		}
	} else if x < tree.Elem {
		tree.Left = add(tree.Left, x)
		if tree.Left.nodeHeight()-tree.Right.nodeHeight() == 2 {
			if x < tree.Left.Elem { //left left
				tree = singleLeftRotate(tree)
			} else { //left right
				tree = doubleRotateLeftRight(tree)
			}
		}
	} else if x > tree.Elem {
		tree.Right = add(tree.Right, x)
		if tree.Right.nodeHeight()-tree.Left.nodeHeight() == 2 {
			if x > tree.Right.Elem {
				tree = singleRightRotate(tree)
			} else {
				tree = doubleRotateRightLeft(tree)
			}
		}
	} else {
		//x existed, we do nothing.
	}

	tree.Height = If(tree.Left.nodeHeight() > tree.Right.nodeHeight(),
		tree.Left.nodeHeight(), tree.Right.nodeHeight()).(int) + 1

	return tree
}

//del an x in AvlTree, invoker format: t = tree.delete(t, x)
func delete(tree *AvlNode, x int) *AvlNode {

	if tree != nil {
		if x < tree.Elem {
			tree.Left = delete(tree.Left, x)
			if tree.Right.nodeHeight()-tree.Left.nodeHeight() == 2 {
				if tree.Right.Right != nil {
					tree = singleRightRotate(tree)
				} else {
					tree = doubleRotateRightLeft(tree)
				}
			} else {
				tree.Height = If(tree.Left.nodeHeight() > tree.Right.nodeHeight(),
					tree.Left.nodeHeight(), tree.Right.nodeHeight()).(int) + 1
			}
		} else if x > tree.Elem {
			tree.Right = delete(tree.Right, x)
			if tree.Left.nodeHeight()-tree.Right.nodeHeight() == 2 {
				if tree.Left.Left != nil {
					tree = singleLeftRotate(tree)
				} else {
					tree = doubleRotateLeftRight(tree)
				}

			} else {
				tree.Height = If(tree.Left.nodeHeight() > tree.Right.nodeHeight(),
					tree.Left.nodeHeight(), tree.Right.nodeHeight()).(int) + 1
			}
		} else if tree.Left != nil && tree.Right != nil {
			tmp := tree.Right.ExtractMin()

			tree.Right = delete(tree.Right, tmp)

		} else {
			if tree.Left == nil {
				tree = tree.Right
			} else if tree.Right == nil {
				tree = tree.Left
			}
		}
	}

	return tree
}

func (tree *AvlNode) Delete(x int) {
	tree = delete(tree, x)
}

/*
//return the max height
func maxHeight(node1, node2 *AvlNode ) int {
	if node1.Height > node2.Height {
		return node1.Height
	} else {
		return node2.Height
	}
}
*/

//left rotate a tree, and update node's height, return the new tree
func singleLeftRotate(tree *AvlNode) *AvlNode {
	var k1 *AvlNode
	if tree != nil {
		k1 = tree.Left
		tree.Left = k1.Right
		k1.Right = tree

		//update height
		tree.Height = If(tree.Left.nodeHeight() > tree.Right.nodeHeight(),
			tree.Left.nodeHeight(), tree.Right.nodeHeight()).(int) + 1
		k1.Height = If(k1.Left.nodeHeight() > k1.Right.nodeHeight(),
			k1.Left.nodeHeight(), k1.Right.nodeHeight()).(int) + 1

		tree = k1
	}

	return tree
}

//right rotate a tree, and update node's height, return the new tree
func singleRightRotate(tree *AvlNode) *AvlNode {
	var k1 *AvlNode
	if tree != nil {
		k1 = tree.Right
		tree.Right = k1.Left
		k1.Left = tree

		//update height
		tree.Height = If(tree.Left.nodeHeight() > tree.Right.nodeHeight(),
			tree.Left.nodeHeight(), tree.Right.nodeHeight()).(int) + 1
		k1.Height = If(k1.Left.nodeHeight() > k1.Right.nodeHeight(),
			k1.Left.nodeHeight(), k1.Right.nodeHeight()).(int) + 1

		tree = k1
	}

	return tree
}

//k3 a tree base, k2: k3's left childe, k1: k2's right childe
//right rotate k2 & k1, left rotate k3 & k3's left child
//return a new tree
func doubleRotateLeftRight(k3 *AvlNode) *AvlNode {
	//right rotatel between k2 & k1
	k3.Left = singleRightRotate(k3.Left)

	//left rotate between k3 and his left child
	return singleLeftRotate(k3)
}

//k3 a tree base, k2: k3's left childe, k1: k2's right childe
//left rotate k2 & k1, right rotate k3 & k3's left child
//return a new tree
func doubleRotateRightLeft(k3 *AvlNode) *AvlNode {
	//left rotatel between k2 & k1
	k3.Right = singleLeftRotate(k3.Right)

	//right rotate between k3 and his left child
	return singleRightRotate(k3)
}
