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
	"math/rand"
	"testing"
)

func nrand(n int) []int {
	i := make([]int, n)
	for ind := range i {
		i[ind] = rand.Int()
	}
	return i
}

func toInterfaces(i []int) []interface{} {
	ifs := make([]interface{}, len(i))
	for ind, v := range i {
		ifs[ind] = v
	}
	return ifs
}

func benchAdd(b *testing.B, s SetGeneric[int]) {
	nums := nrand(b.N)
	b.ResetTimer()
	for _, v := range nums {
		s.Add(v)
	}
}

func BenchmarkAddSafe(b *testing.B) {
	benchAdd(b, NewSetGeneric[int]())
}

func BenchmarkAddUnsafe(b *testing.B) {
	benchAdd(b, NewThreadUnsafeSetGeneric[int]())
}

func benchRemove(b *testing.B, s SetGeneric[int]) {
	nums := nrand(b.N)
	for _, v := range nums {
		s.Add(v)
	}

	b.ResetTimer()
	for _, v := range nums {
		s.Remove(v)
	}
}

func BenchmarkRemoveSafe(b *testing.B) {
	benchRemove(b, NewSetGeneric[int]())
}

func BenchmarkRemoveUnsafe(b *testing.B) {
	benchRemove(b, NewThreadUnsafeSetGeneric[int]())
}

func benchCardinality(b *testing.B, s SetGeneric[int]) {
	for i := 0; i < b.N; i++ {
		s.Cardinality()
	}
}

func BenchmarkCardinalitySafe(b *testing.B) {
	benchCardinality(b, NewSetGeneric[int]())
}

func BenchmarkCardinalityUnsafe(b *testing.B) {
	benchCardinality(b, NewThreadUnsafeSetGeneric[int]())
}

func benchClear(b *testing.B, s SetGeneric[int]) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Clear()
	}
}

func BenchmarkClearSafe(b *testing.B) {
	benchClear(b, NewSetGeneric[int]())
}

func BenchmarkClearUnsafe(b *testing.B) {
	benchClear(b, NewThreadUnsafeSetGeneric[int]())
}

func benchClone(b *testing.B, n int, s SetGeneric[int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Clone()
	}
}

func BenchmarkClone1Safe(b *testing.B) {
	benchClone(b, 1, NewSetGeneric[int]())
}

func BenchmarkClone1Unsafe(b *testing.B) {
	benchClone(b, 1, NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkClone10Safe(b *testing.B) {
	benchClone(b, 10, NewSetGeneric[int]())
}

func BenchmarkClone10Unsafe(b *testing.B) {
	benchClone(b, 10, NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkClone100Safe(b *testing.B) {
	benchClone(b, 100, NewSetGeneric[int]())
}

func BenchmarkClone100Unsafe(b *testing.B) {
	benchClone(b, 100, NewThreadUnsafeSetGeneric[int]())
}

func benchContains(b *testing.B, n int, s SetGeneric[int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(v)
	}

	nums[n-1] = -1 // Definitely not in s

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Contains(nums...)
	}
}

func BenchmarkContains1Safe(b *testing.B) {
	benchContains(b, 1, NewSetGeneric[int]())
}

func BenchmarkContains1Unsafe(b *testing.B) {
	benchContains(b, 1, NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkContains10Safe(b *testing.B) {
	benchContains(b, 10, NewSetGeneric[int]())
}

func BenchmarkContains10Unsafe(b *testing.B) {
	benchContains(b, 10, NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkContains100Safe(b *testing.B) {
	benchContains(b, 100, NewSetGeneric[int]())
}

func BenchmarkContains100Unsafe(b *testing.B) {
	benchContains(b, 100, NewThreadUnsafeSetGeneric[int]())
}

func benchEqual(b *testing.B, n int, s, t SetGeneric[int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(v)
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Equal(t)
	}
}

func BenchmarkEqual1Safe(b *testing.B) {
	benchEqual(b, 1, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkEqual1Unsafe(b *testing.B) {
	benchEqual(b, 1, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkEqual10Safe(b *testing.B) {
	benchEqual(b, 10, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkEqual10Unsafe(b *testing.B) {
	benchEqual(b, 10, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkEqual100Safe(b *testing.B) {
	benchEqual(b, 100, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkEqual100Unsafe(b *testing.B) {
	benchEqual(b, 100, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func benchDifference(b *testing.B, n int, s, t SetGeneric[int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(v)
	}
	for _, v := range nums[:n/2] {
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Difference(t)
	}
}

func benchIsSubset(b *testing.B, n int, s, t SetGeneric[int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(v)
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.IsSubset(t)
	}
}

func BenchmarkIsSubset1Safe(b *testing.B) {
	benchIsSubset(b, 1, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkIsSubset1Unsafe(b *testing.B) {
	benchIsSubset(b, 1, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkIsSubset10Safe(b *testing.B) {
	benchIsSubset(b, 10, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkIsSubset10Unsafe(b *testing.B) {
	benchIsSubset(b, 10, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkIsSubset100Safe(b *testing.B) {
	benchIsSubset(b, 100, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkIsSubset100Unsafe(b *testing.B) {
	benchIsSubset(b, 100, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func benchIsSuperset(b *testing.B, n int, s, t SetGeneric[int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(v)
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.IsSuperset(t)
	}
}

func BenchmarkIsSuperset1Safe(b *testing.B) {
	benchIsSuperset(b, 1, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkIsSuperset1Unsafe(b *testing.B) {
	benchIsSuperset(b, 1, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkIsSuperset10Safe(b *testing.B) {
	benchIsSuperset(b, 10, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkIsSuperset10Unsafe(b *testing.B) {
	benchIsSuperset(b, 10, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkIsSuperset100Safe(b *testing.B) {
	benchIsSuperset(b, 100, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkIsSuperset100Unsafe(b *testing.B) {
	benchIsSuperset(b, 100, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func benchIsProperSubset(b *testing.B, n int, s, t SetGeneric[int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(v)
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.IsProperSubset(t)
	}
}

func BenchmarkIsProperSubset1Safe(b *testing.B) {
	benchIsProperSubset(b, 1, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkIsProperSubset1Unsafe(b *testing.B) {
	benchIsProperSubset(b, 1, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkIsProperSubset10Safe(b *testing.B) {
	benchIsProperSubset(b, 10, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkIsProperSubset10Unsafe(b *testing.B) {
	benchIsProperSubset(b, 10, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkIsProperSubset100Safe(b *testing.B) {
	benchIsProperSubset(b, 100, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkIsProperSubset100Unsafe(b *testing.B) {
	benchIsProperSubset(b, 100, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func benchIsProperSuperset(b *testing.B, n int, s, t SetGeneric[int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(v)
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.IsProperSuperset(t)
	}
}

func BenchmarkIsProperSuperset1Safe(b *testing.B) {
	benchIsProperSuperset(b, 1, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkIsProperSuperset1Unsafe(b *testing.B) {
	benchIsProperSuperset(b, 1, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkIsProperSuperset10Safe(b *testing.B) {
	benchIsProperSuperset(b, 10, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkIsProperSuperset10Unsafe(b *testing.B) {
	benchIsProperSuperset(b, 10, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkIsProperSuperset100Safe(b *testing.B) {
	benchIsProperSuperset(b, 100, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkIsProperSuperset100Unsafe(b *testing.B) {
	benchIsProperSuperset(b, 100, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkDifference1Safe(b *testing.B) {
	benchDifference(b, 1, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkDifference1Unsafe(b *testing.B) {
	benchDifference(b, 1, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkDifference10Safe(b *testing.B) {
	benchDifference(b, 10, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkDifference10Unsafe(b *testing.B) {
	benchDifference(b, 10, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkDifference100Safe(b *testing.B) {
	benchDifference(b, 100, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkDifference100Unsafe(b *testing.B) {
	benchDifference(b, 100, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func benchIntersect(b *testing.B, n int, s, t SetGeneric[int]) {
	nums := nrand(int(float64(n) * float64(1.5)))
	for _, v := range nums[:n] {
		s.Add(v)
	}
	for _, v := range nums[n/2:] {
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Intersect(t)
	}
}

func BenchmarkIntersect1Safe(b *testing.B) {
	benchIntersect(b, 1, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkIntersect1Unsafe(b *testing.B) {
	benchIntersect(b, 1, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkIntersect10Safe(b *testing.B) {
	benchIntersect(b, 10, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkIntersect10Unsafe(b *testing.B) {
	benchIntersect(b, 10, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkIntersect100Safe(b *testing.B) {
	benchIntersect(b, 100, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkIntersect100Unsafe(b *testing.B) {
	benchIntersect(b, 100, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func benchSymmetricDifference(b *testing.B, n int, s, t SetGeneric[int]) {
	nums := nrand(int(float64(n) * float64(1.5)))
	for _, v := range nums[:n] {
		s.Add(v)
	}
	for _, v := range nums[n/2:] {
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.SymmetricDifference(t)
	}
}

func BenchmarkSymmetricDifference1Safe(b *testing.B) {
	benchSymmetricDifference(b, 1, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkSymmetricDifference1Unsafe(b *testing.B) {
	benchSymmetricDifference(b, 1, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkSymmetricDifference10Safe(b *testing.B) {
	benchSymmetricDifference(b, 10, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkSymmetricDifference10Unsafe(b *testing.B) {
	benchSymmetricDifference(b, 10, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkSymmetricDifference100Safe(b *testing.B) {
	benchSymmetricDifference(b, 100, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkSymmetricDifference100Unsafe(b *testing.B) {
	benchSymmetricDifference(b, 100, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func benchUnion(b *testing.B, n int, s, t SetGeneric[int]) {
	nums := nrand(n)
	for _, v := range nums[:n/2] {
		s.Add(v)
	}
	for _, v := range nums[n/2:] {
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Union(t)
	}
}

func BenchmarkUnion1Safe(b *testing.B) {
	benchUnion(b, 1, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkUnion1Unsafe(b *testing.B) {
	benchUnion(b, 1, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkUnion10Safe(b *testing.B) {
	benchUnion(b, 10, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkUnion10Unsafe(b *testing.B) {
	benchUnion(b, 10, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkUnion100Safe(b *testing.B) {
	benchUnion(b, 100, NewSetGeneric[int](), NewSetGeneric[int]())
}

func BenchmarkUnion100Unsafe(b *testing.B) {
	benchUnion(b, 100, NewThreadUnsafeSetGeneric[int](), NewThreadUnsafeSetGeneric[int]())
}

func benchEach(b *testing.B, n int, s SetGeneric[int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Each(func(elem int) bool {
			return false
		})
	}
}

func BenchmarkEach1Safe(b *testing.B) {
	benchEach(b, 1, NewSetGeneric[int]())
}

func BenchmarkEach1Unsafe(b *testing.B) {
	benchEach(b, 1, NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkEach10Safe(b *testing.B) {
	benchEach(b, 10, NewSetGeneric[int]())
}

func BenchmarkEach10Unsafe(b *testing.B) {
	benchEach(b, 10, NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkEach100Safe(b *testing.B) {
	benchEach(b, 100, NewSetGeneric[int]())
}

func BenchmarkEach100Unsafe(b *testing.B) {
	benchEach(b, 100, NewThreadUnsafeSetGeneric[int]())
}

func benchIter(b *testing.B, n int, s SetGeneric[int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := s.Iter()
		for range c {

		}
	}
}

func BenchmarkIter1Safe(b *testing.B) {
	benchIter(b, 1, NewSetGeneric[int]())
}

func BenchmarkIter1Unsafe(b *testing.B) {
	benchIter(b, 1, NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkIter10Safe(b *testing.B) {
	benchIter(b, 10, NewSetGeneric[int]())
}

func BenchmarkIter10Unsafe(b *testing.B) {
	benchIter(b, 10, NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkIter100Safe(b *testing.B) {
	benchIter(b, 100, NewSetGeneric[int]())
}

func BenchmarkIter100Unsafe(b *testing.B) {
	benchIter(b, 100, NewThreadUnsafeSetGeneric[int]())
}

func benchIterator(b *testing.B, n int, s SetGeneric[int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := s.Iterator().C
		for range c {

		}
	}
}

func BenchmarkIterator1Safe(b *testing.B) {
	benchIterator(b, 1, NewSetGeneric[int]())
}

func BenchmarkIterator1Unsafe(b *testing.B) {
	benchIterator(b, 1, NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkIterator10Safe(b *testing.B) {
	benchIterator(b, 10, NewSetGeneric[int]())
}

func BenchmarkIterator10Unsafe(b *testing.B) {
	benchIterator(b, 10, NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkIterator100Safe(b *testing.B) {
	benchIterator(b, 100, NewSetGeneric[int]())
}

func BenchmarkIterator100Unsafe(b *testing.B) {
	benchIterator(b, 100, NewThreadUnsafeSetGeneric[int]())
}

func benchString(b *testing.B, n int, s SetGeneric[int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.String()
	}
}

func BenchmarkString1Safe(b *testing.B) {
	benchString(b, 1, NewSetGeneric[int]())
}

func BenchmarkString1Unsafe(b *testing.B) {
	benchString(b, 1, NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkString10Safe(b *testing.B) {
	benchString(b, 10, NewSetGeneric[int]())
}

func BenchmarkString10Unsafe(b *testing.B) {
	benchString(b, 10, NewThreadUnsafeSetGeneric[int]())
}

func BenchmarkString100Safe(b *testing.B) {
	benchString(b, 100, NewSetGeneric[int]())
}

func BenchmarkString100Unsafe(b *testing.B) {
	benchString(b, 100, NewThreadUnsafeSetGeneric[int]())
}

func benchToSlice(b *testing.B, s SetGeneric[int]) {
	nums := nrand(b.N)
	for _, v := range nums {
		s.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.ToSlice()
	}
}

func BenchmarkToSliceSafe(b *testing.B) {
	benchToSlice(b, NewSetGeneric[int]())
}

func BenchmarkToSliceUnsafe(b *testing.B) {
	benchToSlice(b, NewThreadUnsafeSetGeneric[int]())
}
