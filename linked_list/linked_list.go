package main

import "fmt"

// Node represents an element in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents the singly linked list
type LinkedList struct {
	head *Node
}

// InsertAtFront adds a new node at the beginning
func (list *LinkedList) InsertAtFront(data int) {
	newNode := &Node{data: data}
	newNode.next = list.head
	list.head = newNode
}

// InsertAtBack adds a new node at the end
func (list *LinkedList) InsertAtBack(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// DeleteByValue removes the first node with the given value
func (list *LinkedList) DeleteByValue(value int) {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}
	if list.head.data == value {
		list.head = list.head.next
		return
	}
	current := list.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}
	if current.next == nil {
		fmt.Println("Value not found in list")
		return
	}
	current.next = current.next.next
}

// Search checks if a value exists in the list
func (list *LinkedList) Search(value int) bool {
	current := list.head
	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}

// PrintList prints all elements in the list
func (list *LinkedList) PrintList() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := &LinkedList{}

	list.InsertAtBack(10)
	list.InsertAtBack(20)
	list.InsertAtFront(5)
	list.InsertAtBack(30)

	fmt.Println("Linked List:")
	list.PrintList()

	fmt.Println("Searching for 20:", list.Search(20))
	fmt.Println("Deleting 10")
	list.DeleteByValue(10)
	fmt.Println("Linked List after deletion:")
	list.PrintList()
}
