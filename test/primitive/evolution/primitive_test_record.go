// Code generated by github.com/rogpeppe/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
 */
package avro

import (
	"github.com/rogpeppe/gogen-avro/v7/compiler"
	"github.com/rogpeppe/gogen-avro/v7/vm"
	"github.com/rogpeppe/gogen-avro/v7/vm/types"
	"io"
)

type PrimitiveTestRecord struct {
	LongField int64

	StringField string

	FloatField float32

	BytesField []byte

	DoubleField float64

	IntField int32

	BoolField bool

	NewString string
}

func NewPrimitiveTestRecord() *PrimitiveTestRecord {
	return &PrimitiveTestRecord{}
}

func DeserializePrimitiveTestRecord(r io.Reader) (*PrimitiveTestRecord, error) {
	t := NewPrimitiveTestRecord()
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

func DeserializePrimitiveTestRecordFromSchema(r io.Reader, schema string) (*PrimitiveTestRecord, error) {
	t := NewPrimitiveTestRecord()

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

func writePrimitiveTestRecord(r *PrimitiveTestRecord, w io.Writer) error {
	var err error

	err = vm.WriteLong(r.LongField, w)
	if err != nil {
		return err
	}

	err = vm.WriteString(r.StringField, w)
	if err != nil {
		return err
	}

	err = vm.WriteFloat(r.FloatField, w)
	if err != nil {
		return err
	}

	err = vm.WriteBytes(r.BytesField, w)
	if err != nil {
		return err
	}

	err = vm.WriteDouble(r.DoubleField, w)
	if err != nil {
		return err
	}

	err = vm.WriteInt(r.IntField, w)
	if err != nil {
		return err
	}

	err = vm.WriteBool(r.BoolField, w)
	if err != nil {
		return err
	}

	err = vm.WriteString(r.NewString, w)
	if err != nil {
		return err
	}

	return err
}

func (r *PrimitiveTestRecord) Serialize(w io.Writer) error {
	return writePrimitiveTestRecord(r, w)
}

func (r *PrimitiveTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"LongField\",\"type\":\"long\"},{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"FloatField\",\"type\":\"float\"},{\"name\":\"BytesField\",\"type\":\"bytes\"},{\"name\":\"DoubleField\",\"type\":\"double\"},{\"name\":\"IntField\",\"type\":\"int\"},{\"name\":\"BoolField\",\"type\":\"boolean\"},{\"default\":\"somedefault\",\"name\":\"NewString\",\"type\":\"string\"}],\"name\":\"PrimitiveTestRecord\",\"type\":\"record\"}"
}

func (r *PrimitiveTestRecord) SchemaName() string {
	return "PrimitiveTestRecord"
}

func (_ *PrimitiveTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PrimitiveTestRecord) Get(i int) types.Field {
	switch i {

	case 0:

		return (*types.Long)(&r.LongField)

	case 1:

		return (*types.String)(&r.StringField)

	case 2:

		return (*types.Float)(&r.FloatField)

	case 3:

		return (*types.Bytes)(&r.BytesField)

	case 4:

		return (*types.Double)(&r.DoubleField)

	case 5:

		return (*types.Int)(&r.IntField)

	case 6:

		return (*types.Boolean)(&r.BoolField)

	case 7:

		return (*types.String)(&r.NewString)

	}
	panic("Unknown field index")
}

func (r *PrimitiveTestRecord) SetDefault(i int) {
	switch i {

	case 7:
		r.NewString = "somedefault"
		return

	}
	panic("Unknown field index")
}

func (_ *PrimitiveTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) Finalize()                        {}
