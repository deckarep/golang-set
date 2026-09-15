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

func BenchmarkNew(b *testing.B) {
	for _, c := range []int{1, 10, 100} {
		nums := nrand(c)

		b.Run(buildCaseName(c, true), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = NewSet(nums...)
			}
		})
		b.Run(buildCaseName(c, false), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = NewThreadUnsafeSet(nums...)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(0, c.safe)
		t := nrand(c.n)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, v := range t {
					_ = s.Add(v)
				}
			}
		})
	}
}

func BenchmarkAppend(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(0, c.safe)
		t := nrand(c.n)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.Append(t...)
			}
		})
	}
}

func BenchmarkAppendFrom(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		t := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.AppendFrom(t)
			}
		})
	}
}

func BenchmarkRemove(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		nums := s.ToSlice()

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				n := rand.Intn(c.n)
				s.Clone().Remove(nums[n])
			}
		})
	}
}

func BenchmarkRemoveAll(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		nums := s.ToSlice()

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				n := rand.Intn(c.n)
				s.Clone().RemoveAll(nums[:n]...)
			}
		})
	}
}

func BenchmarkCardinality(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.Cardinality()
			}
		})
	}
}

func BenchmarkClear(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Clone().Clear()
			}
		})
	}
}

func BenchmarkClone(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.Clone()
			}
		})
	}
}

func BenchmarkContainsOne(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		nums := s.ToSlice()

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				n := rand.Intn(c.n)
				_ = s.ContainsOne(nums[n]) // no allocations, using stack-allocated v
			}
		})
	}
}

func BenchmarkContains(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		nums := s.ToSlice()

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				n := rand.Intn(c.n)
				_ = s.Contains(nums[:n]...) // no allocations, using heap-allocated slice
			}
		})
	}
}

func BenchmarkContainsAny(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		nums := s.ToSlice()

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				n := rand.Intn(c.n)
				_ = s.ContainsAny(nums[:n]...) // no allocations, using heap-allocated slice
			}
		})
	}
}

func BenchmarkEqual(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		t := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.Equal(t)
			}
		})
	}
}

func BenchmarkIsSubset(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		t := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.IsSubset(t)
			}
		})
	}
}

func BenchmarkIsSuperset(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		t := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.IsSuperset(t)
			}
		})
	}
}

func BenchmarkIsProperSubset(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		t := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.IsProperSubset(t)
			}
		})
	}
}

func BenchmarkIsProperSuperset(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		t := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.IsProperSuperset(t)
			}
		})
	}
}

func BenchmarkDifference(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		t := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.Difference(t)
			}
		})
	}
}

func BenchmarkIntersect(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		t := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.Intersect(t)
			}
		})
	}
}

func BenchmarkSymmetricDifference(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		t := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.SymmetricDifference(t)
			}
		})
	}
}

func BenchmarkUnion(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		t := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.Union(t)
			}
		})
	}
}

func BenchmarkEach(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Each(func(e int) bool {
					_ = e
					return false
				})
			}
		})
	}
}

func BenchmarkIter(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for e := range s.Iter() {
					_ = e
				}
			}
		})
	}
}

func BenchmarkIterator(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for e := range s.Iterator().C {
					_ = e
				}
			}
		})
	}
}

func BenchmarkString(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.String()
			}
		})
	}
}

func BenchmarkToSlice(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.ToSlice()
			}
		})
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = s.MarshalJSON()
			}
		})
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	for _, c := range buildBenchCases(1, 10, 100) {
		s := buildRandomSet(c.n, c.safe)
		json, _ := s.MarshalJSON()

		b.Run(buildCaseName(c.n, c.safe), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.UnmarshalJSON(json)
			}
		})
	}
}
