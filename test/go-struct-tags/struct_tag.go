// Code generated by github.com/rogpeppe/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     tagstest.avsc
 */
package avro

import (
	"github.com/rogpeppe/gogen-avro/v7/compiler"
	"github.com/rogpeppe/gogen-avro/v7/vm"
	"github.com/rogpeppe/gogen-avro/v7/vm/types"
	"io"
)

type StructTag struct {
	ProductName string `validate:"true"`
}

func NewStructTag() *StructTag {
	return &StructTag{}
}

func DeserializeStructTag(r io.Reader) (*StructTag, error) {
	t := NewStructTag()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeStructTagFromSchema(r io.Reader, schema string) (*StructTag, error) {
	t := NewStructTag()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeStructTag(r *StructTag, w io.Writer) error {
	var err error

	err = vm.WriteString(r.ProductName, w)
	if err != nil {
		return err
	}

	return err
}

func (r *StructTag) Serialize(w io.Writer) error {
	return writeStructTag(r, w)
}

func (r *StructTag) Schema() string {
	return "{\"fields\":[{\"golang.tags\":\"validate:\\\"true\\\"\",\"name\":\"productName\",\"type\":\"string\"}],\"name\":\"StructTag\",\"type\":\"record\"}"
}

func (r *StructTag) SchemaName() string {
	return "StructTag"
}

func (_ *StructTag) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *StructTag) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *StructTag) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *StructTag) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *StructTag) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *StructTag) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *StructTag) SetString(v string)   { panic("Unsupported operation") }
func (_ *StructTag) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *StructTag) Get(i int) types.Field {
	switch i {

	case 0:

		return (*types.String)(&r.ProductName)

	}
	panic("Unknown field index")
}

func (r *StructTag) SetDefault(i int) {
	switch i {

	}
	panic("Unknown field index")
}

func (_ *StructTag) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *StructTag) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *StructTag) Finalize()                        {}
