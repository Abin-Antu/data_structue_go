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
	for index > 0 && h.array[parent(index)] > h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// Pop removes and returns the minimum element (root)
func (h *Minheap) pop() int {
	if len(h.array) == 0 {
		fmt.Println("Heap is empty")
		return -1 // or panic / return error
	}

	min := h.array[0]
	lastIndex := len(h.array) - 1

	// Replace root with last element
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex] // remove last element

	// Heapify down from the root
	h.heapifydown(0)

	return min
}

// Heapify down restores the heap after pop
func (h *Minheap) heapifydown(index int) {
	lastIndex := len(h.array) - 1
	smallest := index

	for {
		left := leftnode(index)
		right := rightnode(index)

		if left <= lastIndex && h.array[left] < h.array[smallest] {
			smallest = left
		}
		if right <= lastIndex && h.array[right] < h.array[smallest] {
			smallest = right
		}

		if smallest != index {
			h.swap(index, smallest)
			index = smallest
		} else {
			break
		}
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

	// Popping elements
	fmt.Println("Popped:", heap.pop())
	fmt.Println("Heap after pop:", heap.array)
}
