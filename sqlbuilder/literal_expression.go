package sqlbuilder

// Representation of an escaped literal
type literalExpression struct {
	expressionInterfaceImpl
	value interface{}
}

func Literal(value interface{}) *literalExpression {
	exp := literalExpression{value: value}
	exp.expressionInterfaceImpl.parent = &exp

	return &exp
}

func (l literalExpression) Serialize(out *queryData, options ...serializeOption) error {
	out.InsertArgument(l.value)

	return nil
}

type numLiteralExpression struct {
	literalExpression
	numericInterfaceImpl
}

func IntLiteral(value int) NumericExpression {
	numLiteral := &numLiteralExpression{}

	numLiteral.literalExpression = *Literal(value)
	numLiteral.literalExpression.parent = numLiteral

	numLiteral.numericInterfaceImpl.parent = numLiteral

	return numLiteral
}
