package mapset

import (
	"fmt"
	"strings"
)

// An OrderedPairGeneric represents a 2-tuple of values.
type OrderedPairGeneric[T comparable] struct {
	First  T
	Second T
}

// Equal says whether two 2-tuples contain the same values in the same order.
func (pair *OrderedPairGeneric[T]) Equal(other OrderedPairGeneric[T]) bool {
	if pair.First == other.First &&
		pair.Second == other.Second {
		return true
	}

	return false
}

// String outputs a 2-tuple in the form "(A, B)".
func (pair OrderedPairGeneric[T]) String() string {
	return fmt.Sprintf("(%v, %v)", pair.First, pair.Second)
}

type threadUnsafeSetGeneric[T comparable] map[T]struct{}

// // An OrderedPair represents a 2-tuple of values.
// type OrderedPair struct {
// 	First  interface{}
// 	Second interface{}
// }

// Assert concrete type:threadUnsafeSetGeneric adheres to SetGeneric interface.
var _ SetGeneric[string] = (*threadUnsafeSetGeneric[string])(nil)

func newThreadUnsafeSetGeneric[T comparable]() threadUnsafeSetGeneric[T] {
	return make(threadUnsafeSetGeneric[T])
}

func (s *threadUnsafeSetGeneric[T]) Add(v T) bool {
	_, found := (*s)[v]
	if found {
		return false //False if it existed already
	}

	(*s)[v] = struct{}{}
	return true
}

func (s *threadUnsafeSetGeneric[T]) Cardinality() int {
	return len(*s)
}

func (s *threadUnsafeSetGeneric[T]) CartesianProduct(other SetGeneric[T]) SetGeneric[any] {
	o := other.(*threadUnsafeSetGeneric[T])
	
	// NOTE: limitation with Go generics or of my knowledge of Go generics?
	
	// I can't seem to declare this without an instantiation cycle.
	//cartProduct := NewThreadUnsafeSetGeneric[OrderedPairGeneric[T]]()
	
	// So here is my crime against humanity.
	cartProduct := NewThreadUnsafeSetGeneric[any]()

	 for i := range *s {
	 	for j := range *o {
			elem := OrderedPairGeneric[T]{First: i, Second: j}
			cartProduct.Add(elem)
	 	}
	 }

	return cartProduct
}

func (s *threadUnsafeSetGeneric[T]) Clear() {
	*s = newThreadUnsafeSetGeneric[T]()
}

func (s *threadUnsafeSetGeneric[T]) Clone() SetGeneric[T] {
	clonedSet := newThreadUnsafeSetGeneric[T]()
	for elem := range *s {
		clonedSet.Add(elem)
	}
	return &clonedSet
}

func (s *threadUnsafeSetGeneric[T]) Contains(v ...T) bool {
	for _, val := range v {
		if _, ok := (*s)[val]; !ok {
			return false
		}
	}
	return true
}

func (s *threadUnsafeSetGeneric[T]) Difference(other SetGeneric[T]) SetGeneric[T] {
	_ = other.(*threadUnsafeSetGeneric[T])

	diff := newThreadUnsafeSetGeneric[T]()
	for elem := range *s {
		if !other.Contains(elem) {
			diff.Add(elem)
		}
	}
	return &diff
}

func (s *threadUnsafeSetGeneric[T]) Each(cb func(T) bool) {
	for elem := range *s {
		if cb(elem) {
			break
		}
	}
}

func (s *threadUnsafeSetGeneric[T]) Equal(other SetGeneric[T]) bool {
	_ = other.(*threadUnsafeSetGeneric[T])

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

func (s *threadUnsafeSetGeneric[T]) Intersect(other SetGeneric[T]) SetGeneric[T] {
	o := other.(*threadUnsafeSetGeneric[T])

	intersection := newThreadUnsafeSetGeneric[T]()
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

func (s *threadUnsafeSetGeneric[T]) IsProperSubset(other SetGeneric[T]) bool {
	return s.IsSubset(other) && !s.Equal(other)
}

func (s *threadUnsafeSetGeneric[T]) IsProperSuperset(other SetGeneric[T]) bool {
	return s.IsSuperset(other) && !s.Equal(other)
}

func (s *threadUnsafeSetGeneric[T]) IsSubset(other SetGeneric[T]) bool {
	_ = other.(*threadUnsafeSetGeneric[T])
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

func (s *threadUnsafeSetGeneric[T]) IsSuperset(other SetGeneric[T]) bool {
	return other.IsSubset(s)
}

func (s *threadUnsafeSetGeneric[T]) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		for elem := range *s {
			ch <- elem
		}
		close(ch)
	}()

	return ch
}

func (s *threadUnsafeSetGeneric[T]) Iterator() *IteratorGeneric[T] {
	iterator, ch, stopCh := newIteratorGeneric[T]()

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

// TODO: how can we make this properly generic, return T but can't return nil.
func (s *threadUnsafeSetGeneric[T]) Pop() any {
	for item := range *s {
		delete(*s, item)
		return item
	}
	return nil
}

// func (s *threadUnsafeSetGeneric[T]) PowerSet() SetGeneric[T] {
// 	powSet := NewThreadUnsafeSet()
// 	nullset := newThreadUnsafeSet()
// 	powSet.Add(&nullset)

// 	for es := range *set {
// 		u := newThreadUnsafeSet()
// 		j := powSet.Iter()
// 		for er := range j {
// 			p := newThreadUnsafeSet()
// 			if reflect.TypeOf(er).Name() == "" {
// 				k := er.(*threadUnsafeSet)
// 				for ek := range *(k) {
// 					p.Add(ek)
// 				}
// 			} else {
// 				p.Add(er)
// 			}
// 			p.Add(es)
// 			u.Add(&p)
// 		}

// 		powSet = powSet.Union(&u)
// 	}

// 	return powSet
// }

func (s *threadUnsafeSetGeneric[T]) Remove(v T) {
	delete(*s, v)
}

func (s *threadUnsafeSetGeneric[T]) String() string {
	items := make([]string, 0, len(*s))

	for elem := range *s {
		items = append(items, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, ", "))
}

func (s *threadUnsafeSetGeneric[T]) SymmetricDifference(other SetGeneric[T]) SetGeneric[T] {
	_ = other.(*threadUnsafeSetGeneric[T])

	a := s.Difference(other)
	b := other.Difference(s)
	return a.Union(b)
}

func (s *threadUnsafeSetGeneric[T]) ToSlice() []T {
	keys := make([]T, 0, s.Cardinality())
	for elem := range *s {
		keys = append(keys, elem)
	}

	return keys
}

func (s *threadUnsafeSetGeneric[T]) Union(other SetGeneric[T]) SetGeneric[T] {
	o := other.(*threadUnsafeSetGeneric[T])

	unionedSet := newThreadUnsafeSetGeneric[T]()

	for elem := range *s {
		unionedSet.Add(elem)
	}
	for elem := range *o {
		unionedSet.Add(elem)
	}
	return &unionedSet
}
