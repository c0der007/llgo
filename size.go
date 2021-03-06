/*
Copyright (c) 2012 Andrew Wilkins <axwalk@gmail.com>

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

package llgo

import (
	"fmt"
	"github.com/axw/llgo/types"
)

func (c *compiler) sizeofType(t types.Type) int {
	switch t := t.(type) {
	case *types.Name:
		return c.sizeofType(t.Underlying)
	case *types.Basic:
		switch t.Kind {
		case types.BoolKind:
			return 1
		case types.IntKind, types.UintKind:
			return 4 // TODO make same as uintptr?
		case types.Int8Kind, types.Uint8Kind:
			return 1
		case types.Int16Kind, types.Uint16Kind:
			return 2
		case types.Int32Kind, types.Uint32Kind, types.Float32Kind:
			return 4
		case types.Int64Kind, types.Uint64Kind, types.Float64Kind, types.Complex64Kind:
			return 8
		case types.Complex128Kind:
			return 16
		case types.StringKind:
			return 4 + c.target.PointerSize()
		case types.UintptrKind, types.UnsafePointerKind:
			return c.target.PointerSize()
		}
	default:
		panic(fmt.Sprintf("unhandled type: %T", t))
	}
	panic("unreachable")
}
