package singlylinkedlist

import (
	"errors"
)

var (
	// ErrIndex is returned when the index is out of the list
	ErrIndex = errors.New("index is out of the list")
	// ErrIndexOf is returned when the index of value can not found
	ErrIndexOf = errors.New("index of value can not found")
)

// List represents a singly linked list structure.
type List struct {
	first *element // the address of the first element
	last  *element // the address of the last element
	size  int      // size of list
}

// element of list.
type element struct {
	value interface{} // can store any type of value
	next  *element    // next stores the address pointing to the next element
}

// New singly linked list.
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) != 0 {
		list.Append(values...)
	}
	return list
}

// List Interface

// Append values (one or more than one) to list.
// i.e. [] -> Append(1, 2, 3) -> [1, 2, 3]
func (list *List) Append(values ...interface{}) {
	if len(values) == 0 {
		return
	}

	// if the size is equal to 0, mean it is a new singly linked list
	if list.size == 0 {
		for i, v := range values {
			newElement := &element{value: v}
			// first element
			if i == 0 {
				list.first = newElement
				list.last = newElement
			} else {
				list.last.next = newElement
				list.last = newElement
			}
			list.size++
		}
	} else {
		for _, v := range values {
			newElement := &element{value: v}
			list.last.next = newElement
			list.last = newElement
			list.size++
		}
	}
}

// PreAppend can append values (one or more than one) to the front of the list.
// i.e. [1, 2] -> Append(3, 4) -> [3, 4, 1, 2]
func (list *List) PreAppend(values ...interface{}) {
	if len(values) == 0 {
		return
	}

	// if the size is equal to 0, mean it is a new singly linked list,
	// Call Append method directly
	if list.size == 0 {
		list.Append(values...)
	} else {
		// reverse traversal values and add to list
		for i := len(values) - 1; i >= 0; i-- {
			newElement := &element{value: values[i]}
			newElement.next = list.first
			list.first = newElement
			list.size++
		}
	}
}

// indexInRange check if the index is within the length of the list.
func (list *List) indexInRange(index int) bool {
	if index >= 0 && index < list.size {
		return true
	}
	return false
}

// Get value by index
func (list *List) Get(index int) (interface{}, error) {
	if !list.indexInRange(index) {
		return nil, ErrIndex
	}

	// find element by index
	foundElement := list.first
	for i := 0; i != index; i++ {
		foundElement = foundElement.next
	}

	return foundElement.value, nil
}

// Remove element by index.
func (list *List) Remove(index int) error {
	if !list.indexInRange(index) {
		return ErrIndex
	}

	foundElement := list.first

	// remove the first element
	if index == 0 {
		list.first = list.first.next
	} else {
		// find element by index
		preFoundElement := new(element)
		for i := 0; i != index; i++ {
			preFoundElement = foundElement
			foundElement = foundElement.next
		}
		// Adjustment pointer
		preFoundElement.next = foundElement.next
	}

	// remove element
	foundElement.value = nil
	foundElement.next = nil
	list.size--

	return nil
}

// Contains returns true if list contains values, false otherwise.
func (list *List) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}

	// TODO

	return false
}

// Swap value by index.
func (list *List) Swap(i, j int) error {
	if !list.indexInRange(i) || !list.indexInRange(j) {
		return ErrIndex
	}
	if i == j {
		return nil
	}

	// find element by i
	foundElementI := list.first
	for index := 0; index != i; index++ {
		foundElementI = foundElementI.next
	}
	// find element by j
	foundElementJ := list.first
	for index := 0; index != j; index++ {
		foundElementJ = foundElementJ.next
	}
	// swap
	foundElementI.value, foundElementJ.value = foundElementJ.value, foundElementI.value

	return nil
}

// Insert value (one or more than one) after index.
func (list *List) Insert(index int, value ...interface{}) error {
	if len(value) == 0 {
		return nil
	}
	if !list.indexInRange(index) {
		return ErrIndex
	}

	// find element by index
	foundElement := list.first
	for i := 0; i != index; i++ {
		foundElement = foundElement.next
	}
	foundElementNext := foundElement.next

	// insert
	for _, v := range value {
		element := &element{value: v}
		foundElement.next = element
		foundElement = element
	}
	foundElement.next = foundElementNext
	list.size += len(value)

	return nil
}

// Set element by index.
func (list *List) Set(index int, value interface{}) error {
	if !list.indexInRange(index) {
		return ErrIndex
	}

	// find element by index
	foundElement := list.first
	for i := 0; i != index; i++ {
		foundElement = foundElement.next
	}
	foundElement.value = value

	return nil
}

// IndexOf get index by value.
func (list *List) IndexOf(value interface{}) (int, error) {
	for i, v := range list.Values() {
		if v == value {
			return i, nil
		}
	}
	return -1, ErrIndexOf
}

// Reverse the list.
func (list *List) Reverse() {
	if list.size == 0 || list.size == 1 {
		return
	}

	// initial assist pointer
	preElement := new(element)
	preElement.value = nil
	preElement.next = nil
	currElement := list.first
	nextElement := list.first.next

	// reset the last pointer
	list.last = currElement

	// reverse
	for currElement != nil {
		currElement.next = preElement
		preElement = currElement
		currElement = nextElement
		if nextElement != nil {
			nextElement = nextElement.next
		}
	}

	// reset the first pointer
	list.first = preElement
}

// Container Interface

// Empty returns true if the list is empty, otherwise returns false.
func (list *List) Empty() bool {
	return list.size == 0
}

// Size returns the size of the list.
func (list *List) Size() int {
	size := list.size
	return size
}

// Clear th list.
func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

// Values returns the values of list.
func (list *List) Values() []interface{} {
	values := make([]interface{}, 0)

	iterator := list.Iterator()
	iterator.Begin()
	for iterator.Next() {
		values = append(values, iterator.Value())
	}

	return values
}
