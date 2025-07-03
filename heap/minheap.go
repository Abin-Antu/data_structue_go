package main

import (
	"fmt"
)

// Minheap defines a min-heap using a slice
type Minheap struct {
	array []int
}

// Get parent index
func parent(i int) int {
	return (i - 1) / 2
}

// Get left child index
func leftnode(i int) int {
	return i*2 + 1
}

// Get right child index
func rightnode(i int) int {
	return i*2 + 2
}

// Swap values in the array
func (h *Minheap) swap(first, second int) {
	h.array[first], h.array[second] = h.array[second], h.array[first]
}

// Insert a value into the heap
func (h *Minheap) insert(item int) {
	h.array = append(h.array, item)
	h.heapifyup(len(h.array) - 1)
}

// Heapify up restores the heap property after insert
func (h *Minheap) heapifyup(index int) {
	// Continue until we reach the root (index 0) or the parent is smaller
	for index > 0 && h.array[parent(index)] > h.array[index] {
		// Swap with the parent
		h.swap(parent(index), index)
		// Move index up to parent
		index = parent(index)
	}
}

func main() {
	heap := &Minheap{}
	heap.insert(10)
	heap.insert(5)
	heap.insert(20)
	heap.insert(3)
	heap.insert(2)
	heap.insert(15)
	fmt.Println("Final Min-Heap:", heap.array)
}
