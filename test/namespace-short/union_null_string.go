// Code generated by github.com/rogpeppe/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"fmt"
	"io"

	"github.com/rogpeppe/gogen-avro/v7/vm"
	"github.com/rogpeppe/gogen-avro/v7/vm/types"
)

type UnionNullStringTypeEnum int

const (
	UnionNullStringTypeEnumNull UnionNullStringTypeEnum = 0

	UnionNullStringTypeEnumString UnionNullStringTypeEnum = 1
)

type UnionNullString struct {
	Null *types.NullVal

	String string

	UnionType UnionNullStringTypeEnum
}

func writeUnionNullString(r *UnionNullString, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {

	case UnionNullStringTypeEnumNull:
		return vm.WriteNull(r.Null, w)

	case UnionNullStringTypeEnumString:
		return vm.WriteString(r.String, w)

	}
	return fmt.Errorf("invalid value for *UnionNullString")
}

func NewUnionNullString() *UnionNullString {
	return &UnionNullString{}
}

func (_ *UnionNullString) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullString) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullString) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullString) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullString) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullString) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionNullString) SetLong(v int64) {
	r.UnionType = (UnionNullStringTypeEnum)(v)
}
func (r *UnionNullString) Get(i int) types.Field {
	switch i {

	case 0:

		return r.Null

	case 1:

		return (*types.String)(&r.String)

	}
	panic("Unknown field index")
}
func (_ *UnionNullString) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullString) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullString) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullString) Finalize()                        {}
