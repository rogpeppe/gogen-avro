// Code generated by github.com/rogpeppe/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     union.avsc
 */
package avro

import (
	"fmt"
	"io"

	"github.com/rogpeppe/gogen-avro/v7/vm"
	"github.com/rogpeppe/gogen-avro/v7/vm/types"
)

type UnionFieldUnionTypeEnum int

const (
	UnionFieldUnionTypeEnumInt UnionFieldUnionTypeEnum = 0

	UnionFieldUnionTypeEnumLong UnionFieldUnionTypeEnum = 1

	UnionFieldUnionTypeEnumFloat UnionFieldUnionTypeEnum = 2

	UnionFieldUnionTypeEnumDouble UnionFieldUnionTypeEnum = 3

	UnionFieldUnionTypeEnumString UnionFieldUnionTypeEnum = 4

	UnionFieldUnionTypeEnumBool UnionFieldUnionTypeEnum = 5

	UnionFieldUnionTypeEnumNull UnionFieldUnionTypeEnum = 6
)

type UnionFieldUnion struct {
	Int int32

	Long int64

	Float float32

	Double float64

	String string

	Bool bool

	Null *types.NullVal

	UnionType UnionFieldUnionTypeEnum
}

func writeUnionFieldUnion(r *UnionFieldUnion, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {

	case UnionFieldUnionTypeEnumInt:
		return vm.WriteInt(r.Int, w)

	case UnionFieldUnionTypeEnumLong:
		return vm.WriteLong(r.Long, w)

	case UnionFieldUnionTypeEnumFloat:
		return vm.WriteFloat(r.Float, w)

	case UnionFieldUnionTypeEnumDouble:
		return vm.WriteDouble(r.Double, w)

	case UnionFieldUnionTypeEnumString:
		return vm.WriteString(r.String, w)

	case UnionFieldUnionTypeEnumBool:
		return vm.WriteBool(r.Bool, w)

	case UnionFieldUnionTypeEnumNull:
		return vm.WriteNull(r.Null, w)

	}
	return fmt.Errorf("invalid value for *UnionFieldUnion")
}

func NewUnionFieldUnion() *UnionFieldUnion {
	return &UnionFieldUnion{}
}

func (_ *UnionFieldUnion) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionFieldUnion) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionFieldUnion) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionFieldUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionFieldUnion) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionFieldUnion) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionFieldUnion) SetLong(v int64) {
	r.UnionType = (UnionFieldUnionTypeEnum)(v)
}
func (r *UnionFieldUnion) Get(i int) types.Field {
	switch i {

	case 0:

		return (*types.Int)(&r.Int)

	case 1:

		return (*types.Long)(&r.Long)

	case 2:

		return (*types.Float)(&r.Float)

	case 3:

		return (*types.Double)(&r.Double)

	case 4:

		return (*types.String)(&r.String)

	case 5:

		return (*types.Boolean)(&r.Bool)

	case 6:

		return r.Null

	}
	panic("Unknown field index")
}
func (_ *UnionFieldUnion) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionFieldUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionFieldUnion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionFieldUnion) Finalize()                        {}
