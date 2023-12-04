package pagination

import (
	"fmt"
)

type FilterOptions struct {
	Fields []Field
}

type Field struct {
	Name     string
	Values   []interface{}
	Operator string
	Type     string
}

func (f *Field) GetQuery() string {
	if f.Operator == "between" {
		switch f.Type {
		case "num":
			return fmt.Sprintf("%s %s %v AND %v", f.Name, f.Operator, f.Values[0], f.Values[1])
		}
		return fmt.Sprintf("%s %s '%v' AND '%v'", f.Name, f.Operator, f.Values[0], f.Values[1])
	}

	switch f.Type {
	case "num":
		return fmt.Sprintf("%s %s %v", f.Name, f.Operator, f.Values[0])
	}
	return fmt.Sprintf("%s %s '%v'", f.Name, f.Operator, f.Values[0])
}

type FieldBuilder struct {
	field *Field
}

func NewFieldBuilder(field *Field) *FieldBuilder {
	return &FieldBuilder{
		field: field,
	}
}

func (fb *FieldBuilder) SetName(name string) {
	fb.field.Name = name
}

func (fb *FieldBuilder) SetOperator(operator string) {
	fb.field.Operator = operator
}

func (fb *FieldBuilder) SetValues(values ...interface{}) {
	fb.field.Values = values
}

func (fb *FieldBuilder) SetType(Type string) {
	fb.field.Type = Type
}

func (fb *FieldBuilder) Build() *Field {
	return fb.field
}

const (
	OperatorEq            = "eq"
	OperatorNotEq         = "neq"
	OperatorLowerThan     = "lt"
	OperatorGreaterThan   = "gt"
	OperatorLowerThanEq   = "lte"
	OperatorGreaterThanEq = "gte"
	OperatorBetween       = ":"
)

func (fo *FilterOptions) AddField(field Field) {
	field.Operator = defineOperator(field.Operator)
	fo.Fields = append(fo.Fields, field)
}

func defineOperator(operator string) string {
	switch operator {
	case OperatorEq:
		return "="
	case OperatorGreaterThan:
		return ">"
	case OperatorLowerThan:
		return "<"
	case OperatorGreaterThanEq:
		return ">="
	case OperatorLowerThanEq:
		return "<="
	case OperatorNotEq:
		return "<>"
	case OperatorBetween:
		return "between"
	default:
		return ""
	}
}
