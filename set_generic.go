/*
Open Source Initiative OSI - The MIT License (MIT):Licensing

The MIT License (MIT)
Copyright (c) 2013 - 2022 Ralph Caraveo (deckarep@gmail.com)

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
// Items stored within it are unordered and unique. It supports
// typical set operations: membership testing, intersection, union,
// difference, symmetric difference and cloning.
//
// Package mapset provides two implementations of the Set
// interface. The default implementation is safe for concurrent
// access, but a non-thread-safe implementation is also provided for
// programs that can benefit from the slight speed improvement and
// that can enforce mutual exclusion through other means.
package mapset

// TODO: delete this...just messing with generics.
// func Equal[T comparable](a, b T) bool {
// 	return a == b
// }

// NewThreadUnsafeSetGeneric creates and returns a reference to an empty set.
// Operations on the resulting set are not thread-safe.
func NewThreadUnsafeSetGeneric[T comparable]() SetGeneric[T] {
	set := newThreadUnsafeSetGeneric[T]()
	return &set
}

// Set is the primary interface provided by the mapset package.  It
// represents an unordered set of data and a large number of
// operations that can be applied to that set.
type SetGeneric[T comparable] interface {
	// Adds an element to the set. Returns whether
	// the item was added.
	Add(val T) bool

	// // Returns the number of elements in the set.
	Cardinality() int

	// // Removes all elements from the set, leaving
	// // the empty set.
	Clear()

	// // Returns a clone of the set using the same
	// // implementation, duplicating all keys.
	Clone() SetGeneric[T]

	// // Returns whether the given items
	// // are all in the set.
	Contains(val ...T) bool

	// // Returns the difference between this set
	// // and other. The returned set will contain
	// // all elements of this set that are not also
	// // elements of other.
	// //
	// // Note that the argument to Difference
	// // must be of the same type as the receiver
	// // of the method. Otherwise, Difference will
	// // panic.
	Difference(other SetGeneric[T]) SetGeneric[T]

	// // Determines if two sets are equal to each
	// // other. If they have the same cardinality
	// // and contain the same elements, they are
	// // considered equal. The order in which
	// // the elements were added is irrelevant.
	// //
	// // Note that the argument to Equal must be
	// // of the same type as the receiver of the
	// // method. Otherwise, Equal will panic.
	Equal(other SetGeneric[T]) bool

	// // Returns a new set containing only the elements
	// // that exist only in both sets.
	// //
	// // Note that the argument to Intersect
	// // must be of the same type as the receiver
	// // of the method. Otherwise, Intersect will
	// // panic.
	Intersect(other SetGeneric[T]) SetGeneric[T]

	// // Determines if every element in this set is in
	// // the other set but the two sets are not equal.
	// //
	// // Note that the argument to IsProperSubset
	// // must be of the same type as the receiver
	// // of the method. Otherwise, IsProperSubset
	// // will panic.
	IsProperSubset(other SetGeneric[T]) bool

	// // Determines if every element in the other set
	// // is in this set but the two sets are not
	// // equal.
	// //
	// // Note that the argument to IsSuperset
	// // must be of the same type as the receiver
	// // of the method. Otherwise, IsSuperset will
	// // panic.
	IsProperSuperset(other SetGeneric[T]) bool

	// // Determines if every element in this set is in
	// // the other set.
	// //
	// // Note that the argument to IsSubset
	// // must be of the same type as the receiver
	// // of the method. Otherwise, IsSubset will
	// // panic.
	IsSubset(other SetGeneric[T]) bool

	// // Determines if every element in the other set
	// // is in this set.
	// //
	// // Note that the argument to IsSuperset
	// // must be of the same type as the receiver
	// // of the method. Otherwise, IsSuperset will
	// // panic.
	IsSuperset(other SetGeneric[T]) bool

	// // Iterates over elements and executes the passed func against each element.
	// // If passed func returns true, stop iteration at the time.
	Each(func(T) bool)

	// // Returns a channel of elements that you can
	// // range over.
	Iter() <-chan T

	// // Returns an Iterator object that you can
	// // use to range over the set.
	Iterator() *IteratorGeneric[T]

	// // Remove a single element from the set.
	Remove(i T)

	// // Provides a convenient string representation
	// // of the current state of the set.
	String() string

	// // Returns a new set with all elements which are
	// // in either this set or the other set but not in both.
	// //
	// // Note that the argument to SymmetricDifference
	// // must be of the same type as the receiver
	// // of the method. Otherwise, SymmetricDifference
	// // will panic.
	SymmetricDifference(other SetGeneric[T]) SetGeneric[T]

	// // Returns a new set with all elements in both sets.
	// //
	// // Note that the argument to Union must be of the

	// // same type as the receiver of the method.
	// // Otherwise, IsSuperset will panic.
	Union(other SetGeneric[T]) SetGeneric[T]

	// // Pop removes and returns an arbitrary item from the set.
	Pop() any

	// Returns all subsets of a given set (Power Set).
	PowerSet() SetGeneric[any]

	// // Returns the Cartesian Product of two sets.
	CartesianProduct(other SetGeneric[T]) SetGeneric[any]

	// // Returns the members of the set as a slice.
	ToSlice() []T
}

// NewSet creates and returns a reference to an empty set.  Operations
// on the resulting set are thread-safe.
// func NewSet(s ...interface{}) Set {
// 	// set := newThreadSafeSet()
// 	// for _, item := range s {
// 	// 	set.Add(item)
// 	// }
// 	// return &set
// 	return nil
// }

// // NewSetWith creates and returns a new set with the given elements.
// // Operations on the resulting set are thread-safe.
// func NewSetWith(elts ...interface{}) Set {
// 	return NewSetFromSlice(elts)
// }

// // NewSetFromSlice creates and returns a reference to a set from an
// // existing slice.  Operations on the resulting set are thread-safe.
// func NewSetFromSlice(s []interface{}) Set {
// 	a := NewSet(s...)
// 	return a
// }

// // NewThreadUnsafeSet creates and returns a reference to an empty set.
// // Operations on the resulting set are not thread-safe.
// func NewThreadUnsafeSet() Set {
// 	set := newThreadUnsafeSet()
// 	return &set
// }

// // NewThreadUnsafeSetFromSlice creates and returns a reference to a
// // set from an existing slice.  Operations on the resulting set are
// // not thread-safe.
// func NewThreadUnsafeSetFromSlice(s []interface{}) Set {
// 	a := NewThreadUnsafeSet()
// 	for _, item := range s {
// 		a.Add(item)
// 	}
// 	return a
// }
