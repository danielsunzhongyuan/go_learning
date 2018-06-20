package field

type boolField struct {
	name      string
	fieldType FieldType
	value     bool
}

func (field *boolField) Name() string {
	return field.name
}

func (field *boolField) Type() FieldType {
	return field.fieldType
}

func (field *boolField) Value() interface{} {
	return field.value
}

func Bool(name string, value bool) Field {
	return &boolField{name: name, fieldType: BoolType, value: value}
}
