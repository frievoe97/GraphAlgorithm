package main

import (
	"fmt"
	"math"
)

type MinHeap struct {
	Elements []*Vertex
}

type MaxHeap struct {
	Elements []*Vertex
}

func (h *MaxHeap) Insert(elem *Vertex) {
	h.Elements = append(h.Elements, elem)

	h.SortMaxHeap(len(h.Elements) - 1)
}

func (h *MinHeap) Insert(elem *Vertex) {
	h.Elements = append(h.Elements, elem)

	h.SortMinHeap(len(h.Elements) - 1)
}

func (h *MaxHeap) ExtractMax() *Vertex {
	firstElement := h.Elements[0]
	h.Elements[0] = h.Elements[len(h.Elements)-1]
	h.Elements = h.Elements[:len(h.Elements)-1]

	h.ExtractMaxChange(0)

	return firstElement
}

func (h *MaxHeap) ExtractMaxChange(index int) {
	indexFirstElement := index
	leftChildIndex := 2*indexFirstElement + 1
	rightChildIndex := 2*indexFirstElement + 2

	// Solange ein Kind kleiner ist, wird das objekt ausgetauscht

	// Both childs exists
	if leftChildIndex < len(h.Elements) && rightChildIndex < len(h.Elements) {

		if compare(h.Elements[leftChildIndex], h.Elements[indexFirstElement]) {

			tempVertex := h.Elements[leftChildIndex]
			h.Elements[leftChildIndex] = h.Elements[indexFirstElement]
			h.Elements[indexFirstElement] = tempVertex

			h.ExtractMaxChange(leftChildIndex)

		} else if compare(h.Elements[rightChildIndex], h.Elements[indexFirstElement]) {
			tempVertex := h.Elements[rightChildIndex]
			h.Elements[rightChildIndex] = h.Elements[indexFirstElement]
			h.Elements[indexFirstElement] = tempVertex

			h.ExtractMaxChange(rightChildIndex)
		}

	} else if leftChildIndex < len(h.Elements) {

		if compare(h.Elements[leftChildIndex], h.Elements[indexFirstElement]) {
			tempVertex := h.Elements[leftChildIndex]
			h.Elements[leftChildIndex] = h.Elements[indexFirstElement]
			h.Elements[indexFirstElement] = tempVertex

			h.ExtractMaxChange(leftChildIndex)
		}

	} else if rightChildIndex < len(h.Elements) {

		if compare(h.Elements[rightChildIndex], h.Elements[indexFirstElement]) {
			tempVertex := h.Elements[rightChildIndex]
			h.Elements[rightChildIndex] = h.Elements[indexFirstElement]
			h.Elements[indexFirstElement] = tempVertex

			h.ExtractMaxChange(rightChildIndex)
		}

	} else {
		return
	}
}

func (h *MinHeap) ExtractMin() *Vertex {
	firstElement := h.Elements[0]
	h.Elements[0] = h.Elements[len(h.Elements)-1]
	h.Elements = h.Elements[:len(h.Elements)-1]

	h.ExtractMinChange(0)

	return firstElement
}

func (h *MinHeap) ExtractMinChange(index int) {
	indexFirstElement := index
	leftChildIndex := 2*indexFirstElement + 1
	rightChildIndex := 2*indexFirstElement + 2

	// Solange ein Kind kleiner ist, wird das objekt ausgetauscht

	// Both childs exists
	if leftChildIndex < len(h.Elements) && rightChildIndex < len(h.Elements) {

		if !compare(h.Elements[leftChildIndex], h.Elements[indexFirstElement]) {

			tempVertex := h.Elements[leftChildIndex]
			h.Elements[leftChildIndex] = h.Elements[indexFirstElement]
			h.Elements[indexFirstElement] = tempVertex

			h.ExtractMinChange(leftChildIndex)

		} else if !compare(h.Elements[rightChildIndex], h.Elements[indexFirstElement]) {
			tempVertex := h.Elements[rightChildIndex]
			h.Elements[rightChildIndex] = h.Elements[indexFirstElement]
			h.Elements[indexFirstElement] = tempVertex

			h.ExtractMinChange(rightChildIndex)
		}

	} else if leftChildIndex < len(h.Elements) {

		if !compare(h.Elements[leftChildIndex], h.Elements[indexFirstElement]) {
			tempVertex := h.Elements[leftChildIndex]
			h.Elements[leftChildIndex] = h.Elements[indexFirstElement]
			h.Elements[indexFirstElement] = tempVertex

			h.ExtractMinChange(leftChildIndex)
		}

	} else if rightChildIndex < len(h.Elements) {

		if !compare(h.Elements[rightChildIndex], h.Elements[indexFirstElement]) {
			tempVertex := h.Elements[rightChildIndex]
			h.Elements[rightChildIndex] = h.Elements[indexFirstElement]
			h.Elements[indexFirstElement] = tempVertex

			h.ExtractMinChange(rightChildIndex)
		}

	} else {
		return
	}
}

func (h *MinHeap) PrintHeap() {
	fmt.Println("--------------------------------- ")
	for _, element := range h.Elements {
		fmt.Printf("Distance: %v\n", element.distance)
	}
}

func (h *MaxHeap) PrintHeap() {
	fmt.Println("--------------------------------- ")
	for _, element := range h.Elements {
		fmt.Printf("Distance: %v\n", element.distance)
	}
}

func (h *MinHeap) SortMinHeap(index int) {

	if index < 2 {
		return
	}

	parentIndex := int(math.Floor(float64(index) / 2.0))

	if compare(h.Elements[parentIndex], h.Elements[index]) {
		tempVertex := h.Elements[parentIndex]
		h.Elements[parentIndex] = h.Elements[index]
		h.Elements[index] = tempVertex
	}

	h.SortMinHeap(parentIndex)
}

func (h *MaxHeap) SortMaxHeap(index int) {

	if index < 2 {
		return
	}

	parentIndex := int(math.Floor(float64(index) / 2.0))

	if compare(h.Elements[index], h.Elements[parentIndex]) {
		tempVertex := h.Elements[parentIndex]
		h.Elements[parentIndex] = h.Elements[index]
		h.Elements[index] = tempVertex
	}

	h.SortMaxHeap(parentIndex)
}
