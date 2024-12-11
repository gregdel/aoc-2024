package aoc

import "fmt"

// List represents a list of elements.
type List[T any] struct {
	Head, Tail *ListElement[T]
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) ForAll(f func(e *ListElement[T])) {
	current := l.Head
	for current != nil {
		f(current)
		current = current.Next
	}
}

func (l *List[T]) Push(e *ListElement[T]) {
	if l.Head == nil {
		l.Head = e
	}
	if l.Tail != nil {
		l.Tail.AddAfter(e)
	}
	l.Tail = e
}

func (l *List[T]) PopHead() *ListElement[T] {
	e := l.Head
	if e == nil {
		return nil
	}

	if e.Next != nil {
		e.Next.Prev = nil
	}

	l.Head = e.Next
	return e
}

func (l *List[T]) Pop() *ListElement[T] {
	e := l.Tail
	if e == nil {
		return nil
	}

	if e.Prev != nil {
		e.Prev.Next = nil
	}

	l.Tail = e.Prev
	return e
}

func (l *List[T]) ReplaceWith(o, n *ListElement[T]) {
	if l.Head == o {
		l.Head = n
	}

	if l.Tail == o {
		l.Tail = n
	}

	o.ReplaceWith(n)
}

// ListElement represents an element of a list.
type ListElement[T any] struct {
	Value      T
	Prev, Next *ListElement[T]
}

func NewListElement[T any](v T) *ListElement[T] {
	return &ListElement[T]{
		Value: v,
	}
}

func (l ListElement[T]) String() string {
	return fmt.Sprintf("%v", l.Value)
}

func (l *ListElement[T]) AddAfter(e *ListElement[T]) {
	next := l.Next
	if next != nil {
		next.Prev = e
	}
	l.Next = e
	e.Prev = l
	e.Next = next
}

func (l *ListElement[T]) AddBefore(e *ListElement[T]) {
	prev := l.Prev
	if prev != nil {
		prev.Next = e
	}
	l.Prev = e
	e.Next = l
	e.Prev = prev
}

func (l *ListElement[T]) ReplaceWith(e *ListElement[T]) {
	prev := l.Prev
	if prev != nil {
		prev.Next = e
	}
	e.Prev = prev

	next := l.Next
	if next != nil {
		next.Prev = e
	}
	e.Next = next
}
