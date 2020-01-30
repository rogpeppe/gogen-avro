// Code generated by github.com/rogpeppe/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     maps.avsc
 */
package avro

import (
	"github.com/rogpeppe/gogen-avro/v7/vm"
	"github.com/rogpeppe/gogen-avro/v7/vm/types"
	"io"
)

func writeMapBytes(r *MapBytes, w io.Writer) error {
	err := vm.WriteLong(int64(len(r.M)), w)
	if err != nil || len(r.M) == 0 {
		return err
	}
	for k, e := range r.M {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = vm.WriteBytes(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapBytes struct {
	keys   []string
	values [][]byte
	M      map[string][]byte
}

func NewMapBytes() *MapBytes {
	return &MapBytes{
		keys:   make([]string, 0),
		values: make([][]byte, 0),
		M:      make(map[string][]byte),
	}
}

func (_ *MapBytes) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapBytes) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapBytes) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapBytes) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapBytes) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapBytes) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapBytes) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapBytes) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapBytes) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapBytes) SetDefault(i int)      { panic("Unsupported operation") }
func (r *MapBytes) Finalize() {
	for i := range r.keys {
		r.M[r.keys[i]] = r.values[i]
	}
	r.keys = nil
	r.values = nil
}

func (r *MapBytes) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v []byte

	r.values = append(r.values, v)

	return (*types.Bytes)(&r.values[len(r.values)-1])

}

func (_ *MapBytes) AppendArray() types.Field { panic("Unsupported operation") }
