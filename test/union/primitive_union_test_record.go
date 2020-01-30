// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     union.avsc
 */
package avro

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

  
type PrimitiveUnionTestRecord struct {

	
	
		UnionField *UnionIntLongFloatDoubleStringBoolNull
	

}

func NewPrimitiveUnionTestRecord() (*PrimitiveUnionTestRecord) {
	return &PrimitiveUnionTestRecord{}
}

func DeserializePrimitiveUnionTestRecord(r io.Reader) (*PrimitiveUnionTestRecord, error) {
	t := NewPrimitiveUnionTestRecord()
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

func DeserializePrimitiveUnionTestRecordFromSchema(r io.Reader, schema string) (*PrimitiveUnionTestRecord, error) {
	t := NewPrimitiveUnionTestRecord()

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

func writePrimitiveUnionTestRecord(r *PrimitiveUnionTestRecord, w io.Writer) error {
	var err error
	
	err = writeUnionIntLongFloatDoubleStringBoolNull( r.UnionField, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *PrimitiveUnionTestRecord) Serialize(w io.Writer) error {
	return writePrimitiveUnionTestRecord(r, w)
}

func (r *PrimitiveUnionTestRecord) Schema() string {
	return "{\"fields\":[{\"default\":1234,\"name\":\"UnionField\",\"type\":[\"int\",\"long\",\"float\",\"double\",\"string\",\"boolean\",\"null\"]}],\"name\":\"PrimitiveUnionTestRecord\",\"type\":\"record\"}"
}

func (r *PrimitiveUnionTestRecord) SchemaName() string {
	return "PrimitiveUnionTestRecord"
}

func (_ *PrimitiveUnionTestRecord) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *PrimitiveUnionTestRecord) SetInt(v int32) { panic("Unsupported operation") }
func (_ *PrimitiveUnionTestRecord) SetLong(v int64) { panic("Unsupported operation") }
func (_ *PrimitiveUnionTestRecord) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *PrimitiveUnionTestRecord) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *PrimitiveUnionTestRecord) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *PrimitiveUnionTestRecord) SetString(v string) { panic("Unsupported operation") }
func (_ *PrimitiveUnionTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PrimitiveUnionTestRecord) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.UnionField = NewUnionIntLongFloatDoubleStringBoolNull()
	
		
		
			return r.UnionField
		
	
	}
	panic("Unknown field index")
}

func (r *PrimitiveUnionTestRecord) SetDefault(i int) {
	switch (i) {
	
        
	case 0:
       	 	r.UnionField = NewUnionIntLongFloatDoubleStringBoolNull()
r.UnionField.Int = 1234
		return
	
	
	}
	panic("Unknown field index")
}

func (_ *PrimitiveUnionTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *PrimitiveUnionTestRecord) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *PrimitiveUnionTestRecord) Finalize() { }
