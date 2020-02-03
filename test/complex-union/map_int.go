// Code generated by github.com/rogpeppe/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     union.avsc
 */
package avro

import (
	"github.com/rogpeppe/gogen-avro/v7/vm"
	"github.com/rogpeppe/gogen-avro/v7/vm/types"
	"io"
)

func writeMapInt(r *MapInt, w io.Writer) error {
	err := vm.WriteLong(int64(len(r.M)), w)
	if err != nil || len(r.M) == 0 {
		return err
	}
	for k, e := range r.M {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = vm.WriteInt(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapInt struct {
	keys   []string
	values []int32
	M      map[string]int32
}

func NewMapInt() *MapInt {
	return &MapInt{
		keys:   make([]string, 0),
		values: make([]int32, 0),
		M:      make(map[string]int32),
	}
}

func (_ *MapInt) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapInt) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapInt) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapInt) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapInt) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapInt) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapInt) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapInt) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapInt) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapInt) SetDefault(i int)      { panic("Unsupported operation") }
func (r *MapInt) Finalize() {
	for i := range r.keys {
		r.M[r.keys[i]] = r.values[i]
	}
	r.keys = nil
	r.values = nil
}

func (r *MapInt) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v int32

	r.values = append(r.values, v)

	return (*types.Int)(&r.values[len(r.values)-1])

}

func (_ *MapInt) AppendArray() types.Field { panic("Unsupported operation") }