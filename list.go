package main

import "fmt"
//LinkedList is a type
type LinkedList struct {
	
	val int
	next *LinkedList}
//SortedLinkedList is a type
type SortedLinkedList struct {

		val int
		next *SortedLinkedList}



type list interface {
	add(int) 
	delete(int)
	display()
	sort()
	extractMax()
	extractMin()

}

func (l *LinkedList) extractMax() int{
	var max int;
	var index int;
	var i int;
	head:=*l
	for l !=nil {
		if max<l.val {
			max=l.val
			index=i
		}
		l=l.next
		i++

	}
	(head).delete(index)
	return max

}

func (l *LinkedList) extractMin() int{
	var min int;
	var index int;
	var i int;
	head:=*l
	for l !=nil {
		if min>l.val {
			min=l.val
			index=i
		}
		l=l.next
		i++

	}
	(head).delete(index)
	return min

}

func (l *LinkedList) add(val int) {
	if (*l).next!=nil{
		(*l).next.add(val)
	}else{
		(*l).next=&LinkedList{val:val,next:nil}
	}
	
}

func (l *LinkedList) delete(index int) {

	if index==0 {
		c:=((*l).next)
		*l=*c
		return
	}
	var n *LinkedList = l
	var prev *LinkedList
	for i := 0; i < index; i++ {
		prev=n
		n= (*n).next
	}
	(*prev).next=(*n).next

	
}
func (l *LinkedList) display() {
	fmt.Println()
	for l !=nil {
		fmt.Printf("%d ",(*l).val)
		l=l.next
	}
}
func findMiddle (l *LinkedList, a **LinkedList, b **LinkedList) {
	slow:=l
	fast:=l.next
	for fast != nil {
		fast=(*fast).next
		if fast!=nil{
			fast=(*fast).next
			slow=(*slow).next
		}
	}
	*a = l; 
    **b = *slow.next; 
    slow.next = nil; 

}


func sort(l *LinkedList) *LinkedList {
	

	if l.next!=nil {
	var a *LinkedList
	var b *LinkedList = &LinkedList{val:1,next:nil}
	findMiddle(l,&a,&b)
	

	firsthalf:=sort(a)
	secondhalf:=sort(b)

	var x *LinkedList = firsthalf;
	var y *LinkedList = secondhalf;
	var merged *LinkedList = &LinkedList{val:0,next:nil};
	head:=merged
	for x!=nil && y!=nil {
		if (*x).val<(*y).val {
			(merged).next=&LinkedList{val:(*x).val,next:nil}
			merged=merged.next
			x=x.next
		}else{
			(merged).next=&LinkedList{val:(*y).val,next:nil}
			merged=merged.next
			y=y.next
		}

	} 
	if (x!=nil){
		merged.next = x

	}else if y!=nil {
		merged.next = y
	}
	
	return (head.next)
} 
return l

 
}
func (l *LinkedList)sort() (*SortedLinkedList){
	*l=*sort(l)

	
	var h *SortedLinkedList =&SortedLinkedList{val:0,next:nil}
	var head *SortedLinkedList=h
	for l!=nil {
		h.next=&SortedLinkedList{val:l.val,next:nil}
		h=h.next
		l=l.next
	}
	return head.next



}



//dhbhjbdbfdsbdhfskbhkdsbksbddbakjdbshbkdjbjksfkjasdhjsfkjdfskhsfdksd



func myfindMiddle (l *SortedLinkedList, a **SortedLinkedList, b **SortedLinkedList) {
	slow:=l
	fast:=l.next
	for fast != nil {
		fast=(*fast).next
		if fast!=nil{
			fast=(*fast).next
			slow=(*slow).next
		}
	}
	*a = l; 
    **b = *slow.next; 
    slow.next = nil; 

}




func mysort(l *SortedLinkedList) *SortedLinkedList {
	

	if l.next!=nil {
	var a *SortedLinkedList
	var b *SortedLinkedList = &SortedLinkedList{val:1,next:nil}
	myfindMiddle(l,&a,&b)
	

	firsthalf:=mysort(a)
	secondhalf:=mysort(b)

	var x *SortedLinkedList = firsthalf;
	var y *SortedLinkedList = secondhalf;
	var merged *SortedLinkedList = &SortedLinkedList{val:0,next:nil};
	head:=merged
	for x!=nil && y!=nil {
		if (*x).val<(*y).val {
			(merged).next=&SortedLinkedList{val:(*x).val,next:nil}
			merged=merged.next
			x=x.next
		}else{
			(merged).next=&SortedLinkedList{val:(*y).val,next:nil}
			merged=merged.next
			y=y.next
		}

	} 
	if (x!=nil){
		merged.next = x

	}else if y!=nil {
		merged.next = y
	}
	
	return (head.next)
} 
return l

 
}
func (l *SortedLinkedList)sort() (*SortedLinkedList){
	*l=*mysort(l)

	
	var h *SortedLinkedList =&SortedLinkedList{val:0,next:nil}
	var head *SortedLinkedList=h
	for l!=nil {
		h.next=&SortedLinkedList{val:l.val,next:nil}
		h=h.next
		l=l.next
	}
	return head.next
}















func (l *SortedLinkedList) add(val int) {
	var prev *SortedLinkedList=l
	for l!=nil && (*l).val<val {
		prev=l
		l=l.next
	}
	
	if l!=nil{
		(*prev).next=&SortedLinkedList{val:val,next:(*prev).next}
	} else {
		(prev).next=&SortedLinkedList{val:val,next:nil}
	}
	
	
}

func (l *SortedLinkedList) delete(index int) {

	if index==0 {
		c:=((*l).next)
		*l=*c
		return
	}
	var n *SortedLinkedList = l
	var prev *SortedLinkedList
	for i := 0; i < index; i++ {
		prev=n
		n= (*n).next
	}
	(*prev).next=(*n).next

	
}
func (l *SortedLinkedList) display() {
	fmt.Println()
	for l !=nil {
		fmt.Printf("%d ",(*l).val)
		l=l.next
	}
}


func (l *SortedLinkedList) extractMax() int{
	var max int;
	var index int;
	var i int;
	head:=*l
	for l !=nil {
		if max<l.val {
			max=l.val
			index=i
		}
		l=l.next
		i++

	}
	(head).delete(index)
	return max

}

func (l *SortedLinkedList) extractMin() int{
	
	min:=(*l).val
	
	
	return min

}