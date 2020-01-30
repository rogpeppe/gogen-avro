// Code generated by github.com/rogpeppe/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"github.com/rogpeppe/gogen-avro/v7/compiler"
	"github.com/rogpeppe/gogen-avro/v7/vm"
	"github.com/rogpeppe/gogen-avro/v7/vm/types"
	"io"
)

// Trace

type HeaderworksTrace struct {

	// Trace Identifier

	TraceId *UnionNullDatatypeUUID
}

func NewHeaderworksTrace() *HeaderworksTrace {
	return &HeaderworksTrace{}
}

func DeserializeHeaderworksTrace(r io.Reader) (*HeaderworksTrace, error) {
	t := NewHeaderworksTrace()
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

func DeserializeHeaderworksTraceFromSchema(r io.Reader, schema string) (*HeaderworksTrace, error) {
	t := NewHeaderworksTrace()

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

func writeHeaderworksTrace(r *HeaderworksTrace, w io.Writer) error {
	var err error

	err = writeUnionNullDatatypeUUID(r.TraceId, w)
	if err != nil {
		return err
	}

	return err
}

func (r *HeaderworksTrace) Serialize(w io.Writer) error {
	return writeHeaderworksTrace(r, w)
}

func (r *HeaderworksTrace) Schema() string {
	return "{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]}],\"name\":\"Trace\",\"type\":\"record\"}"
}

func (r *HeaderworksTrace) SchemaName() string {
	return "headerworks.Trace"
}

func (_ *HeaderworksTrace) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *HeaderworksTrace) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *HeaderworksTrace) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *HeaderworksTrace) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *HeaderworksTrace) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *HeaderworksTrace) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *HeaderworksTrace) SetString(v string)   { panic("Unsupported operation") }
func (_ *HeaderworksTrace) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *HeaderworksTrace) Get(i int) types.Field {
	switch i {

	case 0:

		r.TraceId = NewUnionNullDatatypeUUID()

		return r.TraceId

	}
	panic("Unknown field index")
}

func (r *HeaderworksTrace) SetDefault(i int) {
	switch i {

	case 0:
		r.TraceId = NewUnionNullDatatypeUUID()

		return

	}
	panic("Unknown field index")
}

func (_ *HeaderworksTrace) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *HeaderworksTrace) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *HeaderworksTrace) Finalize()                        {}
