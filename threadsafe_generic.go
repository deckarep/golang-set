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

import "sync"

type threadSafeSetGeneric[T comparable] struct {
	sync.RWMutex
	uss threadUnsafeSetGeneric[T]
}

func newThreadSafeSetGeneric[T comparable]() threadSafeSetGeneric[T] {
	newUss := newThreadUnsafeSetGeneric[T]()
	return threadSafeSetGeneric[T]{
		uss: newUss,
	}
}

func (s *threadSafeSetGeneric[T]) Add(v T) bool {
	s.Lock()
	ret := s.uss.Add(v)
	s.Unlock()
	return ret
}

func (s *threadSafeSetGeneric[T]) Contains(v ...T) bool {
	s.RLock()
	ret := s.uss.Contains(v...)
	s.RUnlock()
	return ret
}

func (s *threadSafeSetGeneric[T]) IsSubset(other SetGeneric[T]) bool {
	o := other.(*threadSafeSetGeneric[T])

	s.RLock()
	o.RLock()

	ret := s.uss.IsSubset(&o.uss)
	s.RUnlock()
	o.RUnlock()
	return ret
}

func (s *threadSafeSetGeneric[T]) IsProperSubset(other SetGeneric[T]) bool {
	o := other.(*threadSafeSetGeneric[T])

	s.RLock()
	defer s.RUnlock()
	o.RLock()
	defer o.RUnlock()

	return s.uss.IsProperSubset(&o.uss)
}

func (s *threadSafeSetGeneric[T]) IsSuperset(other SetGeneric[T]) bool {
	return other.IsSubset(s)
}

func (s *threadSafeSetGeneric[T]) IsProperSuperset(other SetGeneric[T]) bool {
	return other.IsProperSubset(s)
}

func (s *threadSafeSetGeneric[T]) Union(other SetGeneric[T]) SetGeneric[T] {
	o := other.(*threadSafeSetGeneric[T])

	s.RLock()
	o.RLock()

	unsafeUnion := s.uss.Union(&o.uss).(*threadUnsafeSetGeneric[T])
	ret := &threadSafeSetGeneric[T]{uss: *unsafeUnion}
	s.RUnlock()
	o.RUnlock()
	return ret
}

func (s *threadSafeSetGeneric[T]) Intersect(other SetGeneric[T]) SetGeneric[T] {
	o := other.(*threadSafeSetGeneric[T])

	s.RLock()
	o.RLock()

	unsafeIntersection := s.uss.Intersect(&o.uss).(*threadUnsafeSetGeneric[T])
	ret := &threadSafeSetGeneric[T]{uss: *unsafeIntersection}
	s.RUnlock()
	o.RUnlock()
	return ret
}

func (s *threadSafeSetGeneric[T]) Difference(other SetGeneric[T]) SetGeneric[T] {
	o := other.(*threadSafeSetGeneric[T])

	s.RLock()
	o.RLock()

	unsafeDifference := s.uss.Difference(&o.uss).(*threadUnsafeSetGeneric[T])
	ret := &threadSafeSetGeneric[T]{uss: *unsafeDifference}
	s.RUnlock()
	o.RUnlock()
	return ret
}

func (s *threadSafeSetGeneric[T]) SymmetricDifference(other SetGeneric[T]) SetGeneric[T] {
	o := other.(*threadSafeSetGeneric[T])

	s.RLock()
	o.RLock()

	unsafeDifference := s.uss.SymmetricDifference(&o.uss).(*threadUnsafeSetGeneric[T])
	ret := &threadSafeSetGeneric[T]{uss: *unsafeDifference}
	s.RUnlock()
	o.RUnlock()
	return ret
}

func (s *threadSafeSetGeneric[T]) Clear() {
	s.Lock()
	s.uss = newThreadUnsafeSetGeneric[T]()
	s.Unlock()
}

func (s *threadSafeSetGeneric[T]) Remove(v T) {
	s.Lock()
	delete(s.uss, v)
	s.Unlock()
}

func (s *threadSafeSetGeneric[T]) Cardinality() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.uss)
}

func (s *threadSafeSetGeneric[T]) Each(cb func(T) bool) {
	s.RLock()
	for elem := range s.uss {
		if cb(elem) {
			break
		}
	}
	s.RUnlock()
}

func (s *threadSafeSetGeneric[T]) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		s.RLock()

		for elem := range s.uss {
			ch <- elem
		}
		close(ch)
		s.RUnlock()
	}()

	return ch
}

func (s *threadSafeSetGeneric[T]) Iterator() *IteratorGeneric[T] {
 	iterator, ch, stopCh := newIteratorGeneric[T]()

 	go func() {
 		s.RLock()
	L:
 		for elem := range s.uss {
 			select {
 			case <-stopCh:
 				break L
 			case ch <- elem:
 			}
 		}
 		close(ch)
 		s.RUnlock()
 	}()

 	return iterator
}

func (s *threadSafeSetGeneric[T]) Equal(other SetGeneric[T]) bool {
	o := other.(*threadSafeSetGeneric[T])

	s.RLock()
	o.RLock()

	ret := s.uss.Equal(&o.uss)
	s.RUnlock()
	o.RUnlock()
	return ret
}

func (s *threadSafeSetGeneric[T]) Clone() SetGeneric[T] {
	s.RLock()

	unsafeClone := s.uss.Clone().(*threadUnsafeSetGeneric[T])
	ret := &threadSafeSetGeneric[T]{uss: *unsafeClone}
	s.RUnlock()
	return ret
}

func (s *threadSafeSetGeneric[T]) String() string {
	s.RLock()
	ret := s.uss.String()
	s.RUnlock()
	return ret
}

func (s *threadSafeSetGeneric[T]) PowerSet() SetGeneric[any] {
 	s.RLock()
 	unsafePowerSet := s.uss.PowerSet().(*threadUnsafeSetGeneric[any])
 	s.RUnlock()

 	ret := &threadSafeSetGeneric[any]{uss: newThreadUnsafeSetGeneric[any]()}
 	 for subset := range unsafePowerSet.Iter() {
 	 	unsafeSubset := subset.(*threadUnsafeSetGeneric[any])
 	 	ret.Add(&threadSafeSetGeneric[any]{uss: *unsafeSubset})
 	 }
 	return ret
 }

func (s *threadSafeSetGeneric[T]) Pop() any {
	s.Lock()
	defer s.Unlock()
	return s.uss.Pop()
}

func (s *threadSafeSetGeneric[T]) CartesianProduct(other SetGeneric[T]) SetGeneric[any] {
	o := other.(*threadSafeSetGeneric[T])

	s.RLock()
	o.RLock()

	unsafeCartProduct := s.uss.CartesianProduct(&o.uss).(*threadUnsafeSetGeneric[any])
	ret := &threadSafeSetGeneric[any]{uss: *unsafeCartProduct}
	
	s.RUnlock()
	o.RUnlock()
	return ret
}

func (s *threadSafeSetGeneric[T]) ToSlice() []T {
	keys := make([]T, 0, s.Cardinality())
	s.RLock()
	for elem := range s.uss {
		keys = append(keys, elem)
	}
	s.RUnlock()
	return keys
}
