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

package mapset

import "fmt"
import "strings"

type Set struct {
	set map[interface{}]_placeHolder
}

type _placeHolder struct{}

func NewSet() *Set {
	return &Set{make(map[interface{}]_placeHolder)}
}

func (set *Set) Add(i interface{}) bool {
	_, found := set.set[i]
	set.set[i] = _placeHolder{}
	return !found //False if it existed already
}

func (set *Set) Contains(i interface{}) bool {
	if _, found := set.set[i]; found {
		return found //true if it existed already
	}
	return false
}

func (set *Set) IsSubset(other *Set) bool {
	for key, _ := range set.set {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func (set *Set) IsSuperset(other *Set) bool {
	for key, _ := range other.set {
		if !set.Contains(key) {
			return false
		}
	}
	return true
}

func (set *Set) Union(other *Set) *Set {
	if set != nil && other != nil {
		unionedSet := NewSet()

		for key, _ := range set.set {
			unionedSet.Add(key)
		}
		for key, _ := range other.set {
			unionedSet.Add(key)
		}
		return unionedSet
	}
	return nil
}

func (set *Set) Intersect(other *Set) *Set {
	if set != nil && other != nil {
		intersectedSet := NewSet()
		for key, _ := range set.set {
			if other.Contains(key) {
				intersectedSet.Add(key)
			}
		}
		return intersectedSet
	}
	return nil
}

func (set *Set) Difference(other *Set) *Set {
	if set != nil && other != nil {
		differencedSet := NewSet()

		for key, _ := range set.set {
			if !other.Contains(key) {
				differencedSet.Add(key)
			}
		}

		return differencedSet
	}
	return nil
}

func (set *Set) SymmetricDifference(other *Set) *Set {
	if set != nil && other != nil {
		aDiff := set.Difference(other)
		bDiff := other.Difference(set)

		symDifferencedSet := aDiff.Union(bDiff)

		return symDifferencedSet
	}
	return nil
}

func (set *Set) Clear() {
	set.set = make(map[interface{}]_placeHolder)
}

func (set *Set) Remove(i interface{}) {
	delete(set.set, i)
}

func (set *Set) Size() int {
	return len(set.set)
}

func (set *Set) Equal(other *Set) bool {
	if set != nil && other != nil {
		if !(set.Size() == other.Size()) {
			return false
		} else {
			for key, _ := range set.set {
				if !other.Contains(key) {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (set *Set) String() string {
	items := make([]string, 0, len(set.set))

	for key := range set.set {
		items = append(items, fmt.Sprintf("%v", key))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, ", "))
}
