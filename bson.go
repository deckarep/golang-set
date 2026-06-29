//go:build !tinygo

/*
Open Source Initiative OSI - The MIT License (MIT):Licensing

The MIT License (MIT)
Copyright (c) 2013 - 2026 Ralph Caraveo (deckarep@gmail.com)

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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

// bsonMarshaler is embedded in Set[T] (non-tinygo builds) to expose the BSON
// marshal/unmarshal methods as part of the public interface.
type bsonMarshaler interface {
	// MarshalBSONValue will marshal the set into a BSON-based representation.
	MarshalBSONValue() (bsontype.Type, []byte, error)

	// UnmarshalBSONValue will unmarshal a BSON-based byte slice into a full Set datastructure.
	// For this to work, set subtypes must implement the Marshal/Unmarshal interface.
	UnmarshalBSONValue(bt bsontype.Type, b []byte) error
}

// MarshalBSONValue creates a BSON array from the set.
func (s threadUnsafeSet[T]) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(s.ToSlice())
}

// UnmarshalBSONValue recreates a set from a BSON array.
func (s threadUnsafeSet[T]) UnmarshalBSONValue(bt bsontype.Type, b []byte) error {
	if bt != bson.TypeArray {
		return fmt.Errorf("must use BSON Array to unmarshal Set")
	}

	var i []T
	err := bson.UnmarshalValue(bt, b, &i)
	if err != nil {
		return err
	}
	s.append(i...)

	return nil
}

func (t *threadSafeSet[T]) MarshalBSONValue() (bsontype.Type, []byte, error) {
	t.RLock()
	bt, b, err := t.uss.MarshalBSONValue()
	t.RUnlock()

	return bt, b, err
}

func (t *threadSafeSet[T]) UnmarshalBSONValue(bt bsontype.Type, p []byte) error {
	t.Lock()
	err := t.uss.UnmarshalBSONValue(bt, p)
	t.Unlock()

	return err
}
