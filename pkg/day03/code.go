package day03

import (
	"bufio"
	"io"
)

func TaskA(r io.Reader) int {
	var totalPriorities int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		bag := NewBag(line)
		totalPriorities += bag.Priority()
	}
	return totalPriorities
}

func TaskB(r io.Reader) int {
	var totalPriorities int
	scanner := bufio.NewScanner(r)
	for {
		sets := make([]Set[rune], 3)
		for idx := 0; idx < 3; idx++ {
			if !scanner.Scan() {
				return totalPriorities
			}
			set := NewSet[rune]()
			for _, rn := range scanner.Text() {
				set.Add(rn)
			}
			sets[idx] = set
		}
		intersection := sets[0].Intersect(sets[1].Intersect(sets[2]))
		priorities := 0
		for element := range intersection.Elements {
			priorities += computePriority(element)
		}
		totalPriorities += priorities
	}
}

// Set is a naive and incomplete implementation of a generic set data structure.
// It uses a map as its underlying datastructure.
type Set[T comparable] struct {
	Elements map[T]bool
}

// Add adds an element to the Set, by checking whether the value is already present
// in the underlying map as a key.
func (s *Set[T]) Add(elem T) {
	if _, prs := s.Elements[elem]; !prs {
		s.Elements[elem] = true
	}
}

// Intersect identifies the intersection of a set given another set by comparing
// keys in the underlying maps.
func (s Set[T]) Intersect(other Set[T]) Set[T] {
	result := NewSet[T]()
	for k := range s.Elements {
		if _, prs := other.Elements[k]; prs {
			result.Add(k)
		}
	}
	return result
}

func NewSet[T comparable]() Set[T] {
	elements := make(map[T]bool)
	return Set[T]{Elements: elements}
}

// Bag stores the two compartments of an elve's bag with two generic compartments.
type Bag struct {
	LeftCompartment  Set[rune]
	RightCompartment Set[rune]
}

// Priority computes the sum of item priorities by intersecting the left and the
// right compartment of the given bag.
func (b Bag) Priority() int {
	intersection := b.LeftCompartment.Intersect(b.RightCompartment)
	var sum int
	for item := range intersection.Elements {
		sum += computePriority(item)
	}
	return sum
}

func NewBag(input string) *Bag {
	LeftCompartment := NewSet[rune]()
	RightCompartment := NewSet[rune]()
	bag := &Bag{LeftCompartment, RightCompartment}
	for _, rn := range input[:len(input)/2] {
		bag.LeftCompartment.Add(rn)
	}
	for _, rn := range input[len(input)/2:] {
		bag.RightCompartment.Add(rn)
	}
	return bag
}

func computePriority(rn rune) int {
	switch {
	case 'a' <= rn && rn <= 'z':
		return int(rn - 'a' + 1)
	case 'A' <= rn && rn <= 'Z':
		return int(rn - 'A' + 27)
	default:
		return -1
	}
}
