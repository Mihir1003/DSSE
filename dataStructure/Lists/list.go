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
	for (*l).Next != nil {
		l = (*l).Next
	}
	(*l).Next = &LinkedList{Val: Val, Next: nil}

}

func (l *LinkedList) Delete(index int) {

	if index == 0 {
		c := (*l).Next
		*l = *c
		return
	}
	var n = l
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
	fmt.Println()
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
		var b = &LinkedList{Val: 1, Next: nil}
		findMiddle(l, &a, &b)

		firsthalf := Sort(a)
		secondhalf := Sort(b)

		var x = firsthalf
		var y = secondhalf
		var merged = &LinkedList{Val: 0, Next: nil}
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

		return head.Next
	}
	return l

}
func (l *LinkedList) Sort() *SortedLinkedList {
	*l = *Sort(l)

	h := &SortedLinkedList{Val: 0, Next: nil}
	var head = h
	for l != nil {
		h.Next = &SortedLinkedList{Val: l.Val, Next: nil}
		h = h.Next
		l = l.Next
	}
	return head.Next

}
