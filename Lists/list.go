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
	Add(int)
	Delete(int)
	Display()
	Sort()
	ExtractMax()
	ExtractMin()
}

func (l *LinkedList) ExtractMax() int {
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
	(head).Delete(index)
	return max

}

func (l *LinkedList) ExtractMin() int {
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
	(head).Delete(index)
	return min

}

func (l *LinkedList) Add(Val int) {
	if (*l).Next != nil {
		(*l).Next.Add(Val)
	} else {
		(*l).Next = &LinkedList{Val: Val, Next: nil}
	}

}

func (l *LinkedList) Delete(index int) {

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
func (l *LinkedList) Display() {
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

func Sort(l *LinkedList) *LinkedList {

	if l.Next != nil {
		var a *LinkedList
		var b *LinkedList = &LinkedList{Val: 1, Next: nil}
		findMiddle(l, &a, &b)

		firsthalf := Sort(a)
		secondhalf := Sort(b)

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
func (l *LinkedList) Sort() *SortedLinkedList {
	*l = *Sort(l)

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

func mySort(l *SortedLinkedList) *SortedLinkedList {

	if l.Next != nil {
		var a *SortedLinkedList
		var b *SortedLinkedList = &SortedLinkedList{Val: 1, Next: nil}
		myfindMiddle(l, &a, &b)

		firsthalf := mySort(a)
		secondhalf := mySort(b)

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
func (l *SortedLinkedList) Sort() *SortedLinkedList {
	*l = *mySort(l)

	var h *SortedLinkedList = &SortedLinkedList{Val: 0, Next: nil}
	var head *SortedLinkedList = h
	for l != nil {
		h.Next = &SortedLinkedList{Val: l.Val, Next: nil}
		h = h.Next
		l = l.Next
	}
	return head.Next
}

func (l *SortedLinkedList) Add(Val int) {
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

func (l *SortedLinkedList) Delete(index int) {

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
func (l *SortedLinkedList) Display() {
	fmt.Println()
	for l != nil {
		fmt.Printf("%d ", (*l).Val)
		l = l.Next
	}
}

func (l *SortedLinkedList) ExtractMax() int {
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
	(head).Delete(index)
	return max

}

func (l *SortedLinkedList) ExtractMin() int {

	min := (*l).Val

	return min

}
