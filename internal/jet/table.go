package jet

import (
	"github.com/xxxgo/jet/internal/utils"
)

// SerializerTable interface
type SerializerTable interface {
	Serializer
	Table
}

// Table interface
type Table interface {
	columns() []Column
	SchemaName() string
	TableName() string
	AS(alias string)
}

// NewTable creates new table with schema Name, table Name and list of columns
func NewTable(schemaName, name string, columns ...ColumnExpression) SerializerTable {

	t := tableImpl{
		schemaName: schemaName,
		name:       name,
		columnList: columns,
	}

	for _, c := range columns {
		c.setTableName(name)
	}

	return &t
}

type tableImpl struct {
	schemaName string
	name       string
	alias      string
	columnList []ColumnExpression
}

func (t *tableImpl) AS(alias string) {
	t.alias = alias

	for _, c := range t.columnList {
		c.setTableName(alias)
	}
}

func (t *tableImpl) SchemaName() string {
	return t.schemaName
}

func (t *tableImpl) TableName() string {
	return t.name
}

func (t *tableImpl) columns() []Column {
	ret := []Column{}

	for _, col := range t.columnList {
		ret = append(ret, col)
	}

	return ret
}

func (t *tableImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	if t == nil {
		panic("jet: tableImpl is nil")
	}

	out.WriteIdentifier(t.schemaName)
	out.WriteString(".")
	out.WriteIdentifier(t.name)

	if len(t.alias) > 0 {
		out.WriteString("AS")
		out.WriteIdentifier(t.alias)
	}
}

// JoinType is type of table join
type JoinType int

// Table join types
const (
	InnerJoin JoinType = iota
	LeftJoin
	RightJoin
	FullJoin
	CrossJoin
)

// Join expressions are pseudo readable tables.
type joinTableImpl struct {
	lhs         Serializer
	rhs         Serializer
	joinType    JoinType
	onCondition BoolExpression
}

// JoinTable interface
type JoinTable SerializerTable

// NewJoinTable creates new join table
func NewJoinTable(lhs Serializer, rhs Serializer, joinType JoinType, onCondition BoolExpression) JoinTable {

	joinTable := joinTableImpl{
		lhs:         lhs,
		rhs:         rhs,
		joinType:    joinType,
		onCondition: onCondition,
	}

	return &joinTable
}

func (t *joinTableImpl) SchemaName() string {
	if table, ok := t.lhs.(Table); ok {
		return table.SchemaName()
	}
	return ""
}

func (t *joinTableImpl) TableName() string {
	return ""
}

func (t *joinTableImpl) AS(alias string) {
}

func (t *joinTableImpl) columns() []Column {
	var ret []Column

	if lhsTable, ok := t.lhs.(Table); ok {
		ret = append(ret, lhsTable.columns()...)
	}
	if rhsTable, ok := t.rhs.(Table); ok {
		ret = append(ret, rhsTable.columns()...)
	}

	return ret
}

func (t *joinTableImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	if t == nil {
		panic("jet: Join table is nil. ")
	}

	if utils.IsNil(t.lhs) {
		panic("jet: left hand side of join operation is nil table")
	}

	t.lhs.serialize(statement, out)

	out.NewLine()

	switch t.joinType {
	case InnerJoin:
		out.WriteString("INNER JOIN")
	case LeftJoin:
		out.WriteString("LEFT JOIN")
	case RightJoin:
		out.WriteString("RIGHT JOIN")
	case FullJoin:
		out.WriteString("FULL JOIN")
	case CrossJoin:
		out.WriteString("CROSS JOIN")
	}

	if utils.IsNil(t.rhs) {
		panic("jet: right hand side of join operation is nil table")
	}

	t.rhs.serialize(statement, out)

	if t.onCondition == nil && t.joinType != CrossJoin {
		panic("jet: join condition is nil")
	}

	if t.onCondition != nil {
		out.WriteString("ON")
		t.onCondition.serialize(statement, out)
	}
}
