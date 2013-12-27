/*
Open Source Initiative OSI - The MIT License (MIT):Licensing

The MIT License (MIT)
Copyright (c) 2013 Ralph Caraveo (deckarep@gmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Package mapset implements a simple and generic set collection.
// Items stored within it are unordered and unique
// It supports typical set operations: membership testing, intersection, union, difference, symmetric difference and clonning

package mapset

import (
	"fmt"
	"strings"
)

// The primary type that represents a set
type Set map[interface{}]struct{}

// Creates and returns a reference to an empty set.
func NewSet() Set {
	return make(Set)
}

// Creates and returns a reference to a set from an existing slice
func NewSetFromSlice(s []interface{}) Set {
	a := NewSet()
	for _, item := range s {
		a.Add(item)
	}
	return a
}

// Adds an item to the current set if it doesn't already exist in the set.
func (set Set) Add(i interface{}) bool {
	_, found := set[i]
	set[i] = struct{}{}
	return !found //False if it existed already
}

// Determines if a given item is already in the set.
func (set Set) Contains(i interface{}) bool {
	_, found := set[i]
	return found
}

// Determines if the given items are all in the set
func (set Set) ContainsAll(i ...interface{}) bool {
	allSet := NewSetFromSlice(i)
	if allSet.IsSubset(set) {
		return true
	}
	return false
}

// Determines if every item in the other set is in this set.
func (set Set) IsSubset(other Set) bool {
	for elem := range set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Determines if every item of this set is in the other set.
func (set Set) IsSuperset(other Set) bool {
	return other.IsSubset(set)
}

// Returns a new set with all items in both sets.
func (set Set) Union(other Set) Set {
	unionedSet := NewSet()

	for elem := range set {
		unionedSet.Add(elem)
	}
	for elem := range other {
		unionedSet.Add(elem)
	}
	return unionedSet
}

// Returns a new set with items that exist only in both sets.
func (set Set) Intersect(other Set) Set {
	intersection := NewSet()
	// loop over smaller set
	if set.Cardinality() < other.Cardinality() {
		for elem := range set {
			if other.Contains(elem) {
				intersection.Add(elem)
			}
		}
	} else {
		for elem := range other {
			if set.Contains(elem) {
				intersection.Add(elem)
			}
		}
	}
	return intersection
}

// Returns a new set with items in the current set but not in the other set
func (set Set) Difference(other Set) Set {
	differencedSet := NewSet()
	for elem := range set {
		if !other.Contains(elem) {
			differencedSet.Add(elem)
		}
	}
	return differencedSet
}

// Returns a new set with items in the current set or the other set but not in both.
func (set Set) SymmetricDifference(other Set) Set {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clears the entire set to be the empty set.
func (set *Set) Clear() {
	*set = make(Set)
}

// Allows the removal of a single item in the set.
func (set Set) Remove(i interface{}) {
	delete(set, i)
}

// Cardinality returns how many items are currently in the set.
func (set Set) Cardinality() int {
	return len(set)
}

// Iter() returns a channel of type interface{} that you can range over.
func (set Set) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for elem := range set {
			ch <- elem
		}
		close(ch)
	}()

	return ch
}

// Equal determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set Set) Equal(other Set) bool {
	if set.Cardinality() != other.Cardinality() {
		return false
	}
	for elem := range set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Returns a clone of the set.
// Does NOT clone the underlying elements.
func (set Set) Clone() Set {
	clonedSet := NewSet()
	for elem := range set {
		clonedSet.Add(elem)
	}
	return clonedSet
}

// Provides a convenient string representation of the current state of the set.
func (set Set) String() string {
	items := make([]string, 0, len(set))

	for key := range set {
		items = append(items, fmt.Sprintf("%v", key))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, ", "))
}
