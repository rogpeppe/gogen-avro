package schema

// Common methods for all primitive types
type PrimitiveField struct {
	generatorMetadata

	definition interface{}
}

func (s *PrimitiveField) Definition(_ map[QualifiedName]interface{}) (interface{}, error) {
	return s.definition, nil
}

func (s *PrimitiveField) Children() []AvroType {
	return []AvroType{}
}
