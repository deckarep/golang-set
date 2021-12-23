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

package mapset

import (
	"fmt"
	"reflect"
	"strings"
	"encoding/json"
	"bytes"
)

// An OrderedPair represents a 2-tuple of values.
type OrderedPair[T comparable] struct {
	First  T
	Second T
}

// Equal says whether two 2-tuples contain the same values in the same order.
func (pair *OrderedPair[T]) Equal(other OrderedPair[T]) bool {
	if pair.First == other.First &&
		pair.Second == other.Second {
		return true
	}

	return false
}

// String outputs a 2-tuple in the form "(A, B)".
func (pair OrderedPair[T]) String() string {
	return fmt.Sprintf("(%v, %v)", pair.First, pair.Second)
}

type threadUnsafeSet[T comparable] map[T]struct{}

// Assert concrete type:threadUnsafeSet adheres to Set interface.
var _ Set[string] = (*threadUnsafeSet[string])(nil)

func newThreadUnsafeSet[T comparable]() threadUnsafeSet[T] {
	return make(threadUnsafeSet[T])
}

func (s *threadUnsafeSet[T]) Add(v T) bool {
	_, found := (*s)[v]
	if found {
		return false //False if it existed already
	}

	(*s)[v] = struct{}{}
	return true
}

func (s *threadUnsafeSet[T]) Cardinality() int {
	return len(*s)
}

func (s *threadUnsafeSet[T]) CartesianProduct(other Set[T]) Set[any] {
	o := other.(*threadUnsafeSet[T])

	// NOTE: limitation with Go s or of my knowledge of Go s?

	// I can't seem to declare this without an instantiation cycle.
	//cartProduct := NewThreadUnsafeSet[OrderedPair[T]]()

	// So here is my crime against humanity.
	cartProduct := NewThreadUnsafeSet[any]()

	for i := range *s {
		for j := range *o {
			elem := OrderedPair[T]{First: i, Second: j}
			cartProduct.Add(elem)
		}
	}

	return cartProduct
}

func (s *threadUnsafeSet[T]) Clear() {
	*s = newThreadUnsafeSet[T]()
}

func (s *threadUnsafeSet[T]) Clone() Set[T] {
	clonedSet := newThreadUnsafeSet[T]()
	for elem := range *s {
		clonedSet.Add(elem)
	}
	return &clonedSet
}

func (s *threadUnsafeSet[T]) Contains(v ...T) bool {
	for _, val := range v {
		if _, ok := (*s)[val]; !ok {
			return false
		}
	}
	return true
}

func (s *threadUnsafeSet[T]) Difference(other Set[T]) Set[T] {
	_ = other.(*threadUnsafeSet[T])

	diff := newThreadUnsafeSet[T]()
	for elem := range *s {
		if !other.Contains(elem) {
			diff.Add(elem)
		}
	}
	return &diff
}

func (s *threadUnsafeSet[T]) Each(cb func(T) bool) {
	for elem := range *s {
		if cb(elem) {
			break
		}
	}
}

func (s *threadUnsafeSet[T]) Equal(other Set[T]) bool {
	_ = other.(*threadUnsafeSet[T])

	if s.Cardinality() != other.Cardinality() {
		return false
	}
	for elem := range *s {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (s *threadUnsafeSet[T]) Intersect(other Set[T]) Set[T] {
	o := other.(*threadUnsafeSet[T])

	intersection := newThreadUnsafeSet[T]()
	// loop over smaller set
	if s.Cardinality() < other.Cardinality() {
		for elem := range *s {
			if other.Contains(elem) {
				intersection.Add(elem)
			}
		}
	} else {
		for elem := range *o {
			if s.Contains(elem) {
				intersection.Add(elem)
			}
		}
	}
	return &intersection
}

func (s *threadUnsafeSet[T]) IsProperSubset(other Set[T]) bool {
	return s.IsSubset(other) && !s.Equal(other)
}

func (s *threadUnsafeSet[T]) IsProperSuperset(other Set[T]) bool {
	return s.IsSuperset(other) && !s.Equal(other)
}

func (s *threadUnsafeSet[T]) IsSubset(other Set[T]) bool {
	_ = other.(*threadUnsafeSet[T])
	if s.Cardinality() > other.Cardinality() {
		return false
	}
	for elem := range *s {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (s *threadUnsafeSet[T]) IsSuperset(other Set[T]) bool {
	return other.IsSubset(s)
}

func (s *threadUnsafeSet[T]) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		for elem := range *s {
			ch <- elem
		}
		close(ch)
	}()

	return ch
}

func (s *threadUnsafeSet[T]) Iterator() *Iterator[T] {
	iterator, ch, stopCh := newIterator[T]()

	go func() {
	L:
		for elem := range *s {
			select {
			case <-stopCh:
				break L
			case ch <- elem:
			}
		}
		close(ch)
	}()

	return iterator
}

// TODO: how can we make this properly , return T but can't return nil.
func (s *threadUnsafeSet[T]) Pop() any {
	for item := range *s {
		delete(*s, item)
		return item
	}
	return nil
}

func (s *threadUnsafeSet[T]) PowerSet() Set[any] {
	// The type must be any comparable so we have to dumb down to any.
	powSet := NewThreadUnsafeSet[any]()

	nullset := newThreadUnsafeSet[T]()
	powSet.Add(&nullset)

	for es := range *s {
		u := newThreadUnsafeSet[any]()
		j := powSet.Iter()
		for er := range j {
			p := newThreadUnsafeSet[T]()
			if reflect.TypeOf(er).Name() == "" {
				k := er.(*threadUnsafeSet[T])
				for ek := range *(k) {
					p.Add(ek)
				}
			} else {
				p.Add(er.(T))
			}
			p.Add(es)
			u.Add(&p)
		}

		powSet = powSet.Union(&u)
	}

	return powSet
}

func (s *threadUnsafeSet[T]) Remove(v T) {
	delete(*s, v)
}

func (s *threadUnsafeSet[T]) String() string {
	items := make([]string, 0, len(*s))

	for elem := range *s {
		items = append(items, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, ", "))
}

func (s *threadUnsafeSet[T]) SymmetricDifference(other Set[T]) Set[T] {
	_ = other.(*threadUnsafeSet[T])

	a := s.Difference(other)
	b := other.Difference(s)
	return a.Union(b)
}

func (s *threadUnsafeSet[T]) ToSlice() []T {
	keys := make([]T, 0, s.Cardinality())
	for elem := range *s {
		keys = append(keys, elem)
	}

	return keys
}

func (s *threadUnsafeSet[T]) Union(other Set[T]) Set[T] {
	o := other.(*threadUnsafeSet[T])

	unionedSet := newThreadUnsafeSet[T]()

	for elem := range *s {
		unionedSet.Add(elem)
	}
	for elem := range *o {
		unionedSet.Add(elem)
	}
	return &unionedSet
}

// MarshalJSON creates a JSON array from the set, it marshals all elements
func (s *threadUnsafeSet[T]) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, s.Cardinality())

	for elem := range *s {
		b, err := json.Marshal(elem)
		if err != nil {
			return nil, err
		}

		items = append(items, string(b))
	}

	return []byte(fmt.Sprintf("[%s]", strings.Join(items, ","))), nil
}

// UnmarshalJSON recreates a set from a JSON array, it only decodes
// primitive types. Numbers are decoded as json.Number.
func (s *threadUnsafeSet[T]) UnmarshalJSON(b []byte) error {
	var i []any

	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	err := d.Decode(&i)
	if err != nil {
		return err
	}

	for _, v := range i {
		switch t := v.(type) {
		case T:
			s.Add(t)
		default:
			// anything else must be skipped.
			continue
		}
	}

	return nil
}