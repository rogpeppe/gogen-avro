// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)


// A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014
  
type HeaderworksDatatypeUUID struct {

	
	
		Uuid string
	

}

func NewHeaderworksDatatypeUUID() (*HeaderworksDatatypeUUID) {
	return &HeaderworksDatatypeUUID{}
}

func DeserializeHeaderworksDatatypeUUID(r io.Reader) (*HeaderworksDatatypeUUID, error) {
	t := NewHeaderworksDatatypeUUID()
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

func DeserializeHeaderworksDatatypeUUIDFromSchema(r io.Reader, schema string) (*HeaderworksDatatypeUUID, error) {
	t := NewHeaderworksDatatypeUUID()

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

func writeHeaderworksDatatypeUUID(r *HeaderworksDatatypeUUID, w io.Writer) error {
	var err error
	
	err = vm.WriteString( r.Uuid, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *HeaderworksDatatypeUUID) Serialize(w io.Writer) error {
	return writeHeaderworksDatatypeUUID(r, w)
}

func (r *HeaderworksDatatypeUUID) Schema() string {
	return "{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}"
}

func (r *HeaderworksDatatypeUUID) SchemaName() string {
	return "headerworks.datatype.UUID"
}

func (_ *HeaderworksDatatypeUUID) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *HeaderworksDatatypeUUID) SetInt(v int32) { panic("Unsupported operation") }
func (_ *HeaderworksDatatypeUUID) SetLong(v int64) { panic("Unsupported operation") }
func (_ *HeaderworksDatatypeUUID) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *HeaderworksDatatypeUUID) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *HeaderworksDatatypeUUID) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *HeaderworksDatatypeUUID) SetString(v string) { panic("Unsupported operation") }
func (_ *HeaderworksDatatypeUUID) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *HeaderworksDatatypeUUID) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.String)(&r.Uuid)
		
	
	}
	panic("Unknown field index")
}

func (r *HeaderworksDatatypeUUID) SetDefault(i int) {
	switch (i) {
	
        
	case 0:
       	 	r.Uuid = ""
		return
	
	
	}
	panic("Unknown field index")
}

func (_ *HeaderworksDatatypeUUID) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *HeaderworksDatatypeUUID) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *HeaderworksDatatypeUUID) Finalize() { }
