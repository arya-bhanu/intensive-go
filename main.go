package main

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyCollection = errors.New("collection is empty")
)

type Pair[T, U any] struct {
	First  T
	Second U
}

func NewPair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{First: first, Second: second}
}

func (p Pair[T, U]) Swap() Pair[U, T] {
	return Pair[U, T]{
		First:  p.Second,
		Second: p.First,
	}
}

type Stack[T any] struct {
	store []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	stack := s
	if stack == nil {
		stack = NewStack[T]()
	}
	stack.store = append(stack.store, value)
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if s == nil || s.store == nil || len(s.store) == 0 {
		return zero, errors.New("empty stack")
	}
	var copies []T
	var lastValue = s.store[len(s.store)-1]
	for i := 0; i < len(s.store)-1; i++ {
		copies = append(copies, s.store[i])
	}
	s.store = copies
	return lastValue, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s == nil || s.store == nil || len(s.store) == 0 {
		return zero, errors.New("empty stack")
	}
	return s.store[len(s.store)-1], nil
}

func (s *Stack[T]) Size() int {
	stacks := s
	if stacks == nil {
		stacks = &Stack[T]{}
	}

	return len(stacks.store)
}

func (s *Stack[T]) IsEmpty() bool {
	return s == nil || s.store == nil || len(s.store) == 0
}

type Queue[T any] struct {
	store []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(value T) {
	queue := q
	if queue == nil {
		queue = NewQueue[T]()
	}
	queue.store = append(queue.store, value)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q == nil || q.store == nil || len(q.store) == 0 {
		return zero, errors.New("empty queue")
	}
	var copy []T
	front := q.store[0]
	for i := 1; i < len(q.store); i++ {
		copy = append(copy, q.store[i])
	}
	q.store = copy
	return front, nil
}

func (q *Queue[T]) Front() (T, error) {
	var zero T
	if q == nil || q.store == nil || len(q.store) == 0 {
		return zero, errors.New("queue empty")
	}
	return q.store[0], nil
}

func (q *Queue[T]) Size() int {
	queue := q
	if queue == nil {
		queue = &Queue[T]{}
	}

	return len(queue.store)
}

func (q *Queue[T]) IsEmpty() bool {
	return q == nil || q.store == nil || len(q.store) == 0
}

type Set[T comparable] struct {
	store map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	store := make(map[T]bool)
	return &Set[T]{store: store}
}

func (s *Set[T]) Add(value T) {
	sets := s
	if sets == nil {
		sets = NewSet[T]()
	}
	getVal := sets.store[value]
	if !getVal {
		sets.store[value] = true
	}
}

func (s *Set[T]) Remove(value T) {
	if s != nil {
		delete(s.store, value)
	}
}

func (s *Set[T]) Contains(value T) bool {
	if s != nil {
		return s.store[value]
	}
	return false
}

func (s *Set[T]) Size() int {
	if s != nil {
		return len(s.store)
	}
	return 0
}

func (s *Set[T]) Elements() []T {
	var result []T
	if s != nil {
		for key := range s.store {
			result = append(result, key)
		}
	}
	return result
}

func Union[T comparable](s1, s2 *Set[T]) *Set[T] {
	newSet := NewSet[T]()

	if s1 != nil {
		for key := range s1.store {
			newSet.Add(key)
		}
	}
	if s2 != nil {
		for key := range s2.store {
			newSet.Add(key)
		}
	}

	return newSet
}

// Intersection returns a new set containing only elements that exist in both sets
func Intersection[T comparable](s1, s2 *Set[T]) *Set[T] {
	newSet := NewSet[T]()
	if s1 != nil && s2 != nil {
		for key := range s1.store {
			if s2.store[key] {
				newSet.Add(key)
			}
		}
	}
	return newSet
}

// Difference returns a new set with elements in s1 that are not in s2
func Difference[T comparable](s1, s2 *Set[T]) *Set[T] {
	newSet := NewSet[T]()
	if s1 != nil && s2 != nil {
		for key := range s1.store {
			if !s2.store[key] {
				newSet.Add(key)
			}
		}
	}
	return newSet
}

// Filter returns a new slice containing only the elements for which the predicate returns true
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, s := range slice {
		if predicate(s) {
			result = append(result, s)
		}
	}
	return result
}

// Map applies a function to each element in a slice and returns a new slice with the results
func Map[T, U any](slice []T, mapper func(T) U) []U {
	var result []U
	for _, s := range slice {
		result = append(result, mapper(s))
	}
	return result
}

// Reduce reduces a slice to a single value by applying a function to each element
func Reduce[T, U any](slice []T, initial U, reducer func(U, T) U) U {
	for _, s := range slice {
		initial = reducer(initial, s)
	}
	return initial
}

// Contains returns true if the slice contains the given element
func Contains[T comparable](slice []T, element T) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}

// FindIndex returns the index of the first occurrence of the given element or -1 if not found
func FindIndex[T comparable](slice []T, element T) int {
	for i, s := range slice {
		if s == element {
			return i
		}
	}
	return -1
}

// RemoveDuplicates returns a new slice with duplicate elements removed, preserving order
func RemoveDuplicates[T comparable](slice []T) []T {
	sets := NewSet[T]()
	for _, s := range slice {
		sets.Add(s)
	}
	return sets.Elements()
}
func main() {
	fmt.Println("Hello World")
}
