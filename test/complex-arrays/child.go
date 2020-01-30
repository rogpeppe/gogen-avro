// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     arrays.avsc
 */
package avro

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

  
type Child struct {

	
	
		Name string
	

}

func NewChild() (*Child) {
	return &Child{}
}

func DeserializeChild(r io.Reader) (*Child, error) {
	t := NewChild()
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

func DeserializeChildFromSchema(r io.Reader, schema string) (*Child, error) {
	t := NewChild()

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

func writeChild(r *Child, w io.Writer) error {
	var err error
	
	err = vm.WriteString( r.Name, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *Child) Serialize(w io.Writer) error {
	return writeChild(r, w)
}

func (r *Child) Schema() string {
	return "{\"fields\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"Child\",\"type\":\"record\"}"
}

func (r *Child) SchemaName() string {
	return "Child"
}

func (_ *Child) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Child) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Child) SetLong(v int64) { panic("Unsupported operation") }
func (_ *Child) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Child) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Child) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Child) SetString(v string) { panic("Unsupported operation") }
func (_ *Child) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Child) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.String)(&r.Name)
		
	
	}
	panic("Unknown field index")
}

func (r *Child) SetDefault(i int) {
	switch (i) {
	
        
	
	}
	panic("Unknown field index")
}

func (_ *Child) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Child) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Child) Finalize() { }
