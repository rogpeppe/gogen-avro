package namer

import (
	"fmt"

	avro "github.com/actgardner/gogen-avro/schema"
)

var NamerMetadataKey = struct{}{}

type NamerTypeMetadata struct {
	Name              string
	GoType            string
	SerializerMethod  string
	ConstructorMethod string
	WrapperType       string
}

type NamerFieldMetadata struct {
	Name string
}

type Namer struct {
	// typeNamer generates the names of Go types
	typeNamer NameFormatter

	// fieldNamer generates names for struct fields
	fieldNamer NameFormatter
}

func NewNamer(typeNamer NameFormatter, fieldNamer NameFormatter) *Namer {
	return &Namer{
		typeNamer:  typeNamer,
		fieldNamer: fieldNamer,
	}
}

// Apply generates metadata about Go struct names, file paths and method names for generated code
func (n *Namer) Apply(node avro.Node) error {
	// Because the names of parent types depend on child types, names are generated depth-first
	for _, child := range node.Children() {
		err := n.Apply(child)
		if err != nil {
			return err
		}
	}

	metadata := &NamerTypeMetadata{}

	switch t := node.(type) {
	// Traverse References if the underlying definition hasn't been resolved yet
	case *avro.Reference:
		t.SetGeneratorMetadata(NamerMetadataKey, &NamerFieldMetadata{})
		if !t.HasGeneratorMetadata(NamerMetadataKey) {
			n.Apply(t)
		}

	// User-defined types
	case *avro.EnumDefinition:
		metadata.Name = n.typeNamer.Format(t.AvroName().Name)
		metadata.GoType = metadata.Name
		metadata.SerializerMethod = fmt.Sprintf("write%v", metadata.Name)
		metadata.ConstructorMethod = fmt.Sprintf("make(%v, 0)", metadata.GoType)
		metadata.WrapperType = "types.Int"

	case *avro.FixedDefinition:
		metadata.Name = n.typeNamer.Format(t.AvroName().Name)
		metadata.GoType = fmt.Sprintf("*%v", metadata.Name)
		metadata.SerializerMethod = fmt.Sprintf("write%v", metadata.Name)
		metadata.ConstructorMethod = fmt.Sprintf("make(%v, 0)", metadata.GoType)
		metadata.WrapperType = fmt.Sprintf("%vWrapper", metadata.Name)

	case *avro.RecordDefinition:
		fmt.Printf("Namer: %v \nRecordDefinition: %v\n", n, t)
		metadata.Name = n.typeNamer.Format(t.AvroName().Name)
		metadata.GoType = metadata.Name
		metadata.SerializerMethod = fmt.Sprintf("write%v", metadata.Name)
		metadata.ConstructorMethod = fmt.Sprintf("New%v()", metadata.GoType)
		metadata.WrapperType = ""

		for _, f := range t.Fields() {
			f.SetGeneratorMetadata(NamerMetadataKey, &NamerFieldMetadata{n.fieldNamer.Format(f.Name())})
		}

	// Complex types
	case *avro.ArrayField:
		itemMetadata := t.ItemType().GetGeneratorMetadata(NamerMetadataKey).(*NamerTypeMetadata)
		metadata.Name = "Array" + itemMetadata.Name
		metadata.GoType = fmt.Sprintf("[]%v", itemMetadata.GoType)
		metadata.SerializerMethod = fmt.Sprintf("write%v", metadata.Name)
		metadata.ConstructorMethod = fmt.Sprintf("make(%v, 0)", metadata.GoType)
		metadata.WrapperType = fmt.Sprintf("%vWrapper", metadata.Name)

	case *avro.MapField:
		itemMetadata := t.ItemType().GetGeneratorMetadata(NamerMetadataKey).(*NamerTypeMetadata)
		metadata.Name = "Map" + itemMetadata.Name
		metadata.GoType = fmt.Sprintf("*%v", itemMetadata.GoType)
		metadata.SerializerMethod = fmt.Sprintf("write%v", metadata.Name)
		metadata.ConstructorMethod = fmt.Sprintf("New%v()", metadata.Name)
		metadata.WrapperType = ""

	case *avro.UnionField:
		itemMetadata := make([]*NamerTypeMetadata, len(t.AvroTypes()))
		for i, at := range t.AvroTypes() {
			itemMetadata[i] = at.GetGeneratorMetadata(NamerMetadataKey).(*NamerTypeMetadata)
		}
		compositeUnionName := "Union"
		for _, item := range itemMetadata {
			compositeUnionName += item.Name
		}
		metadata.Name = n.typeNamer.Format(compositeUnionName)
		metadata.GoType = "*" + metadata.Name
		metadata.SerializerMethod = fmt.Sprintf("write%v", metadata.Name)
		metadata.ConstructorMethod = fmt.Sprintf("New%v()", metadata.Name)
		metadata.WrapperType = ""

	// Primitive types
	case *avro.BoolField:
		metadata.GoType = "bool"
		metadata.SerializerMethod = "vm.WriteBool"
		metadata.Name = "Bool"
		metadata.WrapperType = "types.Boolean"

	case *avro.BytesField:
		metadata.GoType = "[]byte"
		metadata.SerializerMethod = "vm.WriteBytes"
		metadata.Name = "Bytes"
		metadata.WrapperType = "types.Bytes"

	case *avro.StringField:
		metadata.GoType = "string"
		metadata.SerializerMethod = "vm.WriteString"
		metadata.Name = "String"
		metadata.WrapperType = "types.String"

	case *avro.FloatField:
		metadata.GoType = "float32"
		metadata.SerializerMethod = "vm.WriteFloat"
		metadata.Name = "Float"
		metadata.WrapperType = "types.Float"

	case *avro.DoubleField:
		metadata.GoType = "float64"
		metadata.SerializerMethod = "vm.WriteDouble"
		metadata.Name = "Double"
		metadata.WrapperType = "types.Double"

	case *avro.IntField:
		metadata.GoType = "int32"
		metadata.SerializerMethod = "vm.WriteInt"
		metadata.Name = "Int"
		metadata.WrapperType = "types.Int"

	case *avro.LongField:
		metadata.GoType = "int64"
		metadata.SerializerMethod = "vm.WriteLong"
		metadata.Name = "Long"
		metadata.WrapperType = "types.Long"

	case *avro.NullField:
		metadata.GoType = "*types.NullVal"
		metadata.SerializerMethod = "vm.WriteNull"
		metadata.Name = "Null"
		metadata.WrapperType = ""
	}
	return nil
}
