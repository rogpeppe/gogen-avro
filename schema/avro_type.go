package schema

type AvroType interface {
	Node

	Definition(scope map[QualifiedName]interface{}) (interface{}, error)
	IsReadableBy(f AvroType) bool
}
