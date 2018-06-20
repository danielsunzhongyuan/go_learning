package field

type FieldType int

const (
	UnknownType FieldType = iota
	BoolType
	Int64Type
	Float64Type
	StringType
	ObjectType
)

type Field interface {
	Name() string
	Type() FieldType
	Value() interface{}
}
