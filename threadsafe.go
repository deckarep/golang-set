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

type threadSafeSet[T comparable] struct {
	sync.RWMutex
	uss *threadUnsafeSet[T]
}

func newThreadSafeSet[T comparable]() *threadSafeSet[T] {
	return &threadSafeSet[T]{
		uss: newThreadUnsafeSet[T](),
	}
}

func newThreadSafeSetWithSize[T comparable](cardinality int) *threadSafeSet[T] {
	return &threadSafeSet[T]{
		uss: newThreadUnsafeSetWithSize[T](cardinality),
	}
}

func (t *threadSafeSet[T]) Add(v T) bool {
	t.Lock()
	defer t.Unlock()
	ret := t.uss.Add(v)
	return ret
}

func (t *threadSafeSet[T]) Append(v ...T) int {
	t.Lock()
	defer t.Unlock()
	ret := t.uss.Append(v...)
	return ret
}

func (t *threadSafeSet[T]) Contains(v ...T) bool {
	t.RLock()
	defer t.RUnlock()
	ret := t.uss.Contains(v...)

	return ret
}

func (t *threadSafeSet[T]) ContainsOne(v T) bool {
	t.RLock()
	defer t.RUnlock()
	ret := t.uss.ContainsOne(v)

	return ret
}

func (t *threadSafeSet[T]) ContainsAny(v ...T) bool {
	t.RLock()
	defer t.RUnlock()
	ret := t.uss.ContainsAny(v...)

	return ret
}

func (t *threadSafeSet[T]) ContainsAnyElement(other Set[T]) bool {
	o := other.(*threadSafeSet[T])

	t.RLock()
	o.RLock()
	defer t.RUnlock()
	defer o.RUnlock()

	ret := t.uss.ContainsAnyElement(o.uss)

	return ret
}

func (t *threadSafeSet[T]) IsEmpty() bool {
	return t.Cardinality() == 0
}

func (t *threadSafeSet[T]) IsSubset(other Set[T]) bool {
	o := other.(*threadSafeSet[T])

	t.RLock()
	o.RLock()
	defer t.RUnlock()
	defer o.RUnlock()

	ret := t.uss.IsSubset(o.uss)
	return ret
}

func (t *threadSafeSet[T]) IsProperSubset(other Set[T]) bool {
	o := other.(*threadSafeSet[T])

	t.RLock()
	defer t.RUnlock()
	o.RLock()
	defer o.RUnlock()

	return t.uss.IsProperSubset(o.uss)
}

func (t *threadSafeSet[T]) IsSuperset(other Set[T]) bool {
	return other.IsSubset(t)
}

func (t *threadSafeSet[T]) IsProperSuperset(other Set[T]) bool {
	return other.IsProperSubset(t)
}

func (t *threadSafeSet[T]) Union(other Set[T]) Set[T] {
	o := other.(*threadSafeSet[T])

	t.RLock()
	o.RLock()
	defer t.RUnlock()
	defer o.RUnlock()

	unsafeUnion := t.uss.Union(o.uss).(*threadUnsafeSet[T])
	ret := &threadSafeSet[T]{uss: unsafeUnion}
	return ret
}

func (t *threadSafeSet[T]) Intersect(other Set[T]) Set[T] {
	o := other.(*threadSafeSet[T])

	t.RLock()
	o.RLock()
	defer t.RUnlock()
	defer o.RUnlock()

	unsafeIntersection := t.uss.Intersect(o.uss).(*threadUnsafeSet[T])
	ret := &threadSafeSet[T]{uss: unsafeIntersection}
	return ret
}

func (t *threadSafeSet[T]) Difference(other Set[T]) Set[T] {
	o := other.(*threadSafeSet[T])

	t.RLock()
	o.RLock()
	defer t.RUnlock()
	defer o.RUnlock()

	unsafeDifference := t.uss.Difference(o.uss).(*threadUnsafeSet[T])
	ret := &threadSafeSet[T]{uss: unsafeDifference}
	return ret
}

func (t *threadSafeSet[T]) SymmetricDifference(other Set[T]) Set[T] {
	o := other.(*threadSafeSet[T])

	t.RLock()
	o.RLock()
	defer t.RUnlock()
	defer o.RUnlock()

	unsafeDifference := t.uss.SymmetricDifference(o.uss).(*threadUnsafeSet[T])
	ret := &threadSafeSet[T]{uss: unsafeDifference}
	return ret
}

func (t *threadSafeSet[T]) Clear() {
	t.Lock()
	defer t.Unlock()
	t.uss.Clear()
}

func (t *threadSafeSet[T]) Remove(v T) {
	t.Lock()
	defer t.Unlock()
	delete(*t.uss, v)
}

func (t *threadSafeSet[T]) RemoveAll(i ...T) {
	t.Lock()
	defer t.Unlock()
	t.uss.RemoveAll(i...)
}

func (t *threadSafeSet[T]) Cardinality() int {
	t.RLock()
	defer t.RUnlock()
	return len(*t.uss)
}

func (t *threadSafeSet[T]) Each(cb func(T) bool) {
	t.RLock()
	defer t.RUnlock()
	for elem := range *t.uss {
		if cb(elem) {
			break
		}
	}
}

func (t *threadSafeSet[T]) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		t.RLock()
		defer t.RUnlock()
		for elem := range *t.uss {
			ch <- elem
		}
		close(ch)
	}()

	return ch
}

func (t *threadSafeSet[T]) Iterator() *Iterator[T] {
	iterator, ch, stopCh := newIterator[T]()

	go func() {
		t.RLock()
		defer t.RUnlock()
	L:
		for elem := range *t.uss {
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

func (t *threadSafeSet[T]) Equal(other Set[T]) bool {
	o := other.(*threadSafeSet[T])

	t.RLock()
	o.RLock()
	defer t.RUnlock()
	defer o.RUnlock()

	ret := t.uss.Equal(o.uss)
	return ret
}

func (t *threadSafeSet[T]) Clone() Set[T] {
	t.RLock()
	defer t.RUnlock()

	unsafeClone := t.uss.Clone().(*threadUnsafeSet[T])
	ret := &threadSafeSet[T]{uss: unsafeClone}
	return ret
}

func (t *threadSafeSet[T]) String() string {
	t.RLock()
	defer t.RUnlock()
	ret := t.uss.String()
	return ret
}

func (t *threadSafeSet[T]) Pop() (T, bool) {
	t.Lock()
	defer t.Unlock()
	return t.uss.Pop()
}

func (t *threadSafeSet[T]) ToSlice() []T {
	keys := make([]T, 0, t.Cardinality())
	t.RLock()
	defer t.RUnlock()
	for elem := range *t.uss {
		keys = append(keys, elem)
	}
	return keys
}

func (t *threadSafeSet[T]) MarshalJSON() ([]byte, error) {
	t.RLock()
	defer t.RUnlock()
	b, err := t.uss.MarshalJSON()

	return b, err
}

func (t *threadSafeSet[T]) UnmarshalJSON(p []byte) error {
	t.RLock()
	defer t.RUnlock()
	err := t.uss.UnmarshalJSON(p)

	return err
}
