package Lists

import "fmt"

//LinkedList is a type
type LinkedList struct {
	Val  int
	Next *LinkedList
}

//SortedLinkedList is a type
type SortedLinkedList struct {
	Val  int
	Next *SortedLinkedList
}

type list interface {
	add(int)
	delete(int)
	display()
	sort()
	extractMax()
	extractMin()
}

func (l *LinkedList) extractMax() int {
	var max int
	var index int
	var i int
	head := *l
	for l != nil {
		if max < l.Val {
			max = l.Val
			index = i
		}
		l = l.Next
		i++

	}
	(head).delete(index)
	return max

}

func (l *LinkedList) extractMin() int {
	var min int
	var index int
	var i int
	head := *l
	for l != nil {
		if min > l.Val {
			min = l.Val
			index = i
		}
		l = l.Next
		i++

	}
	(head).delete(index)
	return min

}

func (l *LinkedList) add(Val int) {
	if (*l).Next != nil {
		(*l).Next.add(Val)
	} else {
		(*l).Next = &LinkedList{Val: Val, Next: nil}
	}

}

func (l *LinkedList) delete(index int) {

	if index == 0 {
		c := ((*l).Next)
		*l = *c
		return
	}
	var n *LinkedList = l
	var prev *LinkedList
	for i := 0; i < index; i++ {
		prev = n
		n = (*n).Next
	}
	(*prev).Next = (*n).Next

}
func (l *LinkedList) display() {
	fmt.Println()
	for l != nil {
		fmt.Printf("%d ", (*l).Val)
		l = l.Next
	}
}
func findMiddle(l *LinkedList, a **LinkedList, b **LinkedList) {
	slow := l
	fast := l.Next
	for fast != nil {
		fast = (*fast).Next
		if fast != nil {
			fast = (*fast).Next
			slow = (*slow).Next
		}
	}
	*a = l
	**b = *slow.Next
	slow.Next = nil

}

func sort(l *LinkedList) *LinkedList {

	if l.Next != nil {
		var a *LinkedList
		var b *LinkedList = &LinkedList{Val: 1, Next: nil}
		findMiddle(l, &a, &b)

		firsthalf := sort(a)
		secondhalf := sort(b)

		var x *LinkedList = firsthalf
		var y *LinkedList = secondhalf
		var merged *LinkedList = &LinkedList{Val: 0, Next: nil}
		head := merged
		for x != nil && y != nil {
			if (*x).Val < (*y).Val {
				(merged).Next = &LinkedList{Val: (*x).Val, Next: nil}
				merged = merged.Next
				x = x.Next
			} else {
				(merged).Next = &LinkedList{Val: (*y).Val, Next: nil}
				merged = merged.Next
				y = y.Next
			}

		}
		if x != nil {
			merged.Next = x

		} else if y != nil {
			merged.Next = y
		}

		return (head.Next)
	}
	return l

}
func (l *LinkedList) sort() *SortedLinkedList {
	*l = *sort(l)

	var h *SortedLinkedList = &SortedLinkedList{Val: 0, Next: nil}
	var head *SortedLinkedList = h
	for l != nil {
		h.Next = &SortedLinkedList{Val: l.Val, Next: nil}
		h = h.Next
		l = l.Next
	}
	return head.Next

}

//dhbhjbdbfdsbdhfskbhkdsbksbddbakjdbshbkdjbjksfkjasdhjsfkjdfskhsfdksd

func myfindMiddle(l *SortedLinkedList, a **SortedLinkedList, b **SortedLinkedList) {
	slow := l
	fast := l.Next
	for fast != nil {
		fast = (*fast).Next
		if fast != nil {
			fast = (*fast).Next
			slow = (*slow).Next
		}
	}
	*a = l
	**b = *slow.Next
	slow.Next = nil

}

func mysort(l *SortedLinkedList) *SortedLinkedList {

	if l.Next != nil {
		var a *SortedLinkedList
		var b *SortedLinkedList = &SortedLinkedList{Val: 1, Next: nil}
		myfindMiddle(l, &a, &b)

		firsthalf := mysort(a)
		secondhalf := mysort(b)

		var x *SortedLinkedList = firsthalf
		var y *SortedLinkedList = secondhalf
		var merged *SortedLinkedList = &SortedLinkedList{Val: 0, Next: nil}
		head := merged
		for x != nil && y != nil {
			if (*x).Val < (*y).Val {
				(merged).Next = &SortedLinkedList{Val: (*x).Val, Next: nil}
				merged = merged.Next
				x = x.Next
			} else {
				(merged).Next = &SortedLinkedList{Val: (*y).Val, Next: nil}
				merged = merged.Next
				y = y.Next
			}

		}
		if x != nil {
			merged.Next = x

		} else if y != nil {
			merged.Next = y
		}

		return (head.Next)
	}
	return l

}
func (l *SortedLinkedList) sort() *SortedLinkedList {
	*l = *mysort(l)

	var h *SortedLinkedList = &SortedLinkedList{Val: 0, Next: nil}
	var head *SortedLinkedList = h
	for l != nil {
		h.Next = &SortedLinkedList{Val: l.Val, Next: nil}
		h = h.Next
		l = l.Next
	}
	return head.Next
}

func (l *SortedLinkedList) add(Val int) {
	var prev *SortedLinkedList = l
	for l != nil && (*l).Val < Val {
		prev = l
		l = l.Next
	}

	if l != nil {
		(*prev).Next = &SortedLinkedList{Val: Val, Next: (*prev).Next}
	} else {
		(prev).Next = &SortedLinkedList{Val: Val, Next: nil}
	}

}

func (l *SortedLinkedList) delete(index int) {

	if index == 0 {
		c := ((*l).Next)
		*l = *c
		return
	}
	var n *SortedLinkedList = l
	var prev *SortedLinkedList
	for i := 0; i < index; i++ {
		prev = n
		n = (*n).Next
	}
	(*prev).Next = (*n).Next

}
func (l *SortedLinkedList) display() {
	fmt.Println()
	for l != nil {
		fmt.Printf("%d ", (*l).Val)
		l = l.Next
	}
}

func (l *SortedLinkedList) extractMax() int {
	var max int
	var index int
	var i int
	head := *l
	for l != nil {
		if max < l.Val {
			max = l.Val
			index = i
		}
		l = l.Next
		i++

	}
	(head).delete(index)
	return max

}

func (l *SortedLinkedList) extractMin() int {

	min := (*l).Val

	return min

}
