package main

import "fmt"

type Node struct {
	val int
	next * Node
}

type Stack struct {
	head * Node
}

func(st * Stack) push(val int) {
	newNode := &Node{val: val, next:nil}
	if newNode == nil {
		st.head = newNode
	} else {
		newNode.next = st.head
		st.head = newNode
	}
}

func (st * Stack) pop() {
	st.head = st.head.next
}

func (st * Stack) peek() int {
	return st.head.val
}

func (st * Stack) clear() {
	st.head = nil
}

func (st * Stack) contains (n int) bool{
	cur := st.head
	for cur != nil{
		if cur.val == n {
			return true
		}
		cur = cur.next
	}
	return false
}

func (st * Stack) increment (){
	cur := st.head
	for cur != nil {
		cur.val += 1
		cur = cur.next
	}
}

func (st*Stack) print() {
	cur := st.head
	for cur != nil {
		fmt.Print(cur.val, " ")
		cur = cur.next
	}
	fmt.Println()
}

func (st* Stack) print_rev(){
	cur := st.head
	arr := make([]int, 0)
	for cur != nil {
		arr = append(arr, cur.val)
		cur = cur.next
	}
  
	for i := len(arr) - 1; i >= 0; i-- {
	  	fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}


func main() {
	var s Stack
	s.push(3)
	s.push(7)
	s.push(8)
	s.peek()
	s.pop()
	s.contains(1)
	s.print()
	s.increment()
	s.print_rev()
	s.clear()
}