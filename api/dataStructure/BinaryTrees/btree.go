package BinaryTrees

// Item the type of the binary search tree

// Node a single node that composes the tree
type Node struct {
	Key   int
	Left  *Node //Left
	Right *Node //Right
}

// ItemBinarySearchTree the binary search tree of Items
type ItemBinarySearchTree struct {
	Root *Node
}

// Insert inserts the Item t in the tree
func (bst *ItemBinarySearchTree) Add(Key int) {

	n := &Node{Key, nil, nil}
	if bst.Root == nil {
		bst.Root = n
	} else {
		insertNode(bst.Root, n)
	}
}

// internal function to find the correct place for a node in a tree
func insertNode(node, newNode *Node) {
	if newNode.Key < node.Key {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertNode(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertNode(node.Right, newNode)
		}
	}
}

// Min returns the Item with min value stored in the tree
func (bst *ItemBinarySearchTree) ExtractMin() int {

	n := bst.Root
	if n == nil {
		return 0
	}
	for {
		if n.Left == nil {
			return n.Key
		}
		n = n.Left
	}
}

// Max returns the Item with max value stored in the tree
func (bst *ItemBinarySearchTree) ExtractMax() int {
	n := bst.Root
	if n == nil {
		return 0
	}
	for {
		if n.Right == nil {
			return n.Key
		}
		n = n.Right
	}
}

// Search returns true if the Item t exists in the tree
func (bst *ItemBinarySearchTree) Search(Key int) bool {
	return search(bst.Root, Key)
}

// internal recursive function to search an item in the tree
func search(n *Node, Key int) bool {
	if n == nil {
		return false
	}
	if Key < n.Key {
		return search(n.Left, Key)
	}
	if Key > n.Key {
		return search(n.Right, Key)
	}
	return true
}

// Remove removes the Item with Key `Key` from the tree
func (bst *ItemBinarySearchTree) Delete(Key int) {
	remove(bst.Root, Key)
}

// internal recursive function to remove an item
func remove(node *Node, Key int) *Node {
	if node == nil {
		return nil
	}
	if Key < node.Key {
		node.Left = remove(node.Left, Key)
		return node
	}
	if Key > node.Key {
		node.Right = remove(node.Right, Key)
		return node
	}
	// Key == node.Key
	if node.Left == nil && node.Right == nil {
		node = nil
		return nil
	}
	if node.Left == nil {
		node = node.Right
		return node
	}
	if node.Right == nil {
		node = node.Left
		return node
	}
	leftmostrightside := node.Right
	for {
		//find smallest value on the Right side
		if leftmostrightside != nil && leftmostrightside.Left != nil {
			leftmostrightside = leftmostrightside.Left
		} else {
			break
		}
	}
	node.Key = leftmostrightside.Key
	node.Right = remove(node.Right, node.Key)
	return node
}

// String prints a visual representation of the tree
