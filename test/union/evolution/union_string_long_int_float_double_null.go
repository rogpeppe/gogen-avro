// Code generated by github.com/rogpeppe/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
 */
package avro

import (
	"fmt"
	"io"

	"github.com/rogpeppe/gogen-avro/v7/vm"
	"github.com/rogpeppe/gogen-avro/v7/vm/types"
)

type UnionStringLongIntFloatDoubleNullTypeEnum int

const (
	UnionStringLongIntFloatDoubleNullTypeEnumString UnionStringLongIntFloatDoubleNullTypeEnum = 0

	UnionStringLongIntFloatDoubleNullTypeEnumLong UnionStringLongIntFloatDoubleNullTypeEnum = 1

	UnionStringLongIntFloatDoubleNullTypeEnumInt UnionStringLongIntFloatDoubleNullTypeEnum = 2

	UnionStringLongIntFloatDoubleNullTypeEnumFloat UnionStringLongIntFloatDoubleNullTypeEnum = 3

	UnionStringLongIntFloatDoubleNullTypeEnumDouble UnionStringLongIntFloatDoubleNullTypeEnum = 4

	UnionStringLongIntFloatDoubleNullTypeEnumNull UnionStringLongIntFloatDoubleNullTypeEnum = 5
)

type UnionStringLongIntFloatDoubleNull struct {
	String string

	Long int64

	Int int32

	Float float32

	Double float64

	Null *types.NullVal

	UnionType UnionStringLongIntFloatDoubleNullTypeEnum
}

func writeUnionStringLongIntFloatDoubleNull(r *UnionStringLongIntFloatDoubleNull, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {

	case UnionStringLongIntFloatDoubleNullTypeEnumString:
		return vm.WriteString(r.String, w)

	case UnionStringLongIntFloatDoubleNullTypeEnumLong:
		return vm.WriteLong(r.Long, w)

	case UnionStringLongIntFloatDoubleNullTypeEnumInt:
		return vm.WriteInt(r.Int, w)

	case UnionStringLongIntFloatDoubleNullTypeEnumFloat:
		return vm.WriteFloat(r.Float, w)

	case UnionStringLongIntFloatDoubleNullTypeEnumDouble:
		return vm.WriteDouble(r.Double, w)

	case UnionStringLongIntFloatDoubleNullTypeEnumNull:
		return vm.WriteNull(r.Null, w)

	}
	return fmt.Errorf("invalid value for *UnionStringLongIntFloatDoubleNull")
}

func NewUnionStringLongIntFloatDoubleNull() *UnionStringLongIntFloatDoubleNull {
	return &UnionStringLongIntFloatDoubleNull{}
}

func (_ *UnionStringLongIntFloatDoubleNull) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNull) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNull) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNull) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNull) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNull) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionStringLongIntFloatDoubleNull) SetLong(v int64) {
	r.UnionType = (UnionStringLongIntFloatDoubleNullTypeEnum)(v)
}
func (r *UnionStringLongIntFloatDoubleNull) Get(i int) types.Field {
	switch i {

	case 0:

		return (*types.String)(&r.String)

	case 1:

		return (*types.Long)(&r.Long)

	case 2:

		return (*types.Int)(&r.Int)

	case 3:

		return (*types.Float)(&r.Float)

	case 4:

		return (*types.Double)(&r.Double)

	case 5:

		return r.Null

	}
	panic("Unknown field index")
}
func (_ *UnionStringLongIntFloatDoubleNull) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNull) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionStringLongIntFloatDoubleNull) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNull) Finalize()                {}
